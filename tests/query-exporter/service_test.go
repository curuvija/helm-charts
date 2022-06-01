package query_exporter

import (
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/require"
	networkingv1 "k8s.io/api/networking/v1"
	"path/filepath"
	"testing"
)

func TestIngressEnabledCreatesIngress(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs(filepath.Join("..", "..", "charts", "query-exporter"))
	require.NoError(t, err)

	options := &helm.Options{
		ValuesFiles: []string{
			filepath.Join(helmChartPath, "values.yaml"),
		},
		SetValues: map[string]string{"ingress.enabled": "true"},
	}

	rendered, err := helm.RenderTemplateE(t, options, helmChartPath, "ingress", []string{"templates/ingress.yaml"})
	var ingress networkingv1.Ingress
	helm.UnmarshalK8SYaml(t, rendered, &ingress)

	ingressSpecRules := ingress.Spec.Rules
	require.Equal(t, ingressSpecRules[0].Host, "chart-example.local")
	require.Equal(t, ingressSpecRules[0].HTTP.Paths[0].Path, "/")
	require.Equal(t, ingressSpecRules[0].HTTP.Paths[0].Backend.Service.Port.Number, int32(9560))
	require.NotEmpty(t, rendered)
	require.NoError(t, err)

}

func TestDefaultValuesIngressNotEnabledDoesNotCreateIngress(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs(filepath.Join("..", "..", "charts", "query-exporter"))
	require.NoError(t, err)

	options := &helm.Options{
		ValuesFiles: []string{
			filepath.Join(helmChartPath, "values.yaml"),
		},
		SetValues: map[string]string{"ingress.enabled": "false"},
	}

	_, err = helm.RenderTemplateE(t, options, helmChartPath, "ingress", []string{"templates/ingress.yaml"})
	expected := "error while running command: exit status 1; Error: could not find template templates/ingress.yaml in chart"
	require.EqualError(t, err, expected)

}
