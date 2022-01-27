package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	csite "github.com/cloudwan/workplace-sdk/client/v1alpha2/site"
	rsite "github.com/cloudwan/workplace-sdk/resources/v1alpha2/site"

	"github.com/cloudwan/edgelq-sdk/examples/utils"
)

/* This is simple illustration of how to construct a client and send a request.
 *
 * Based on provided parameters (endpoint, credentials, projectID) it lists Site resources
 * from Workplace service.
 *
 * In order to use execute example, you need either user account (prepare access token, like one from web browser),
 * or credentials for service account (JSON format, as defined as in edgelq-sdk/common/api/credentials.proto -
 * ServiceAccount). Of course, user or service account will need permission to list probes in project you
 * select, so this is pre-requirement.
 */

func main() {
	ctx := context.Background()

	// ------------------------- Get params and just verify they were provided -------------------------
	endpoint := pflag.StringP("endpoint", "e", "", "API endpoint (for example, workplace.stg01b.edgelq.com:443)")
	projectId := pflag.StringP("project", "p", "", "Name of the project to list sites from")
	accessToken := pflag.StringP("access-token", "a", "", "Active token for your user")
	credsFile := pflag.StringP("credentials", "c", "", "Path to service account credentials file")
	pflag.Parse()
	if *projectId == "" {
		pflag.Usage()
		panic(fmt.Errorf("no project ID was provided"))
	}
	if *endpoint == "" {
		pflag.Usage()
		panic(fmt.Errorf("no endpoint was provided"))
	}
	if *accessToken == "" && *credsFile == "" {
		pflag.Usage()
		panic(fmt.Errorf("access token OR credentials file are necessary in order to execute this example"))
	}

	// ------------------------- Create GRPC connection -------------------------
	grpcConn := utils.Dial(ctx, *endpoint, *accessToken, *credsFile)
	siteClient := csite.NewSiteServiceClient(grpcConn)

	// ------------------------- Lets begin! -------------------------
	listSitesResp, err := siteClient.ListSites(ctx, &csite.ListSitesRequest{
		Parent: rsite.NewNameBuilder().SetProjectId(*projectId).SetRegionId(gotenresource.WildcardId).Parent(),
		View:   view.View_BASIC,
	})
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("failed to list sites: %s\n", err))
		os.Exit(1)
	}

	for _, site := range listSitesResp.GetSites() {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("received site %s\n\n", site))
	}
}
