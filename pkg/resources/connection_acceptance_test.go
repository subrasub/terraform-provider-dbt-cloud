package resources_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/gthesheep/terraform-provider-dbt-cloud/pkg/dbt_cloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDbtCloudConnectionResource(t *testing.T) {

	connectionName := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	connectionName2 := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	projectName := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	oAuthClientID := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	oAuthClientSecret := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDbtCloudConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDbtCloudConnectionResourceBasicConfig(connectionName, projectName, oAuthClientID, oAuthClientSecret),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbtCloudConnectionExists("dbt_cloud_connection.test_connection"),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_connection", "name", connectionName),
				),
			},
			// RENAME
			{
				Config: testAccDbtCloudConnectionResourceBasicConfig(connectionName2, projectName, oAuthClientID, oAuthClientSecret),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbtCloudConnectionExists("dbt_cloud_connection.test_connection"),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_connection", "name", connectionName2),
				),
			},
			// 			// MODIFY
			// 			{
			// 				Config: testAccDbtCloudEnvironmentResourceModifiedConfig(projectName, projectName2, environmentName2),
			// 				Check: resource.ComposeTestCheckFunc(
			// 					testAccCheckDbtCloudProjectExists("dbt_cloud_environment.test_env"),
			// 					resource.TestCheckResourceAttr("dbt_cloud_environment.test_env", "name", environmentName2),
			// 					resource.TestCheckResourceAttr("dbt_cloud_environment.test_env", "dbt_version", "1.0.1"),
			// 				),
			// 			},
			// IMPORT
			{
				ResourceName:            "dbt_cloud_connection.test_connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"oauth_client_id", "oauth_client_secret"},
			},
		},
	})
}

func TestAccDbtCloudRedshiftConnectionResource(t *testing.T) {

	connectionName := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	connectionName2 := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	projectName := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDbtCloudConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDbtCloudConnectionResourceRedshiftConfig(connectionName, projectName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbtCloudConnectionExists("dbt_cloud_connection.test_redshift_connection"),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_redshift_connection", "name", connectionName),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_redshift_connection", "database", "db"),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_redshift_connection", "port", "5432"),
				),
			},
			// RENAME
			{
				Config: testAccDbtCloudConnectionResourceRedshiftConfig(connectionName2, projectName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbtCloudConnectionExists("dbt_cloud_connection.test_redshift_connection"),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_redshift_connection", "name", connectionName2),
				),
			},
			// 			// MODIFY
			// 			{
			// 				Config: testAccDbtCloudEnvironmentResourceModifiedConfig(projectName, projectName2, environmentName2),
			// 				Check: resource.ComposeTestCheckFunc(
			// 					testAccCheckDbtCloudProjectExists("dbt_cloud_environment.test_env"),
			// 					resource.TestCheckResourceAttr("dbt_cloud_environment.test_env", "name", environmentName2),
			// 					resource.TestCheckResourceAttr("dbt_cloud_environment.test_env", "dbt_version", "1.0.1"),
			// 				),
			// 			},
			// IMPORT
			{
				ResourceName:            "dbt_cloud_connection.test_redshift_connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"oauth_client_id", "oauth_client_secret"},
			},
		},
	})
}

func TestAccDbtCloudPostgresConnectionResource(t *testing.T) {

	connectionName := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	connectionName2 := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	projectName := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDbtCloudConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDbtCloudConnectionResourcePostgresConfig(connectionName, projectName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbtCloudConnectionExists("dbt_cloud_connection.test_postgres_connection"),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_postgres_connection", "name", connectionName),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_postgres_connection", "database", "db"),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_postgres_connection", "port", "5432"),
				),
			},
			// RENAME
			{
				Config: testAccDbtCloudConnectionResourcePostgresConfig(connectionName2, projectName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbtCloudConnectionExists("dbt_cloud_connection.test_postgres_connection"),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_postgres_connection", "name", connectionName2),
				),
			},
			// 			// MODIFY
			// 			{
			// 				Config: testAccDbtCloudEnvironmentResourceModifiedConfig(projectName, projectName2, environmentName2),
			// 				Check: resource.ComposeTestCheckFunc(
			// 					testAccCheckDbtCloudProjectExists("dbt_cloud_environment.test_env"),
			// 					resource.TestCheckResourceAttr("dbt_cloud_environment.test_env", "name", environmentName2),
			// 					resource.TestCheckResourceAttr("dbt_cloud_environment.test_env", "dbt_version", "1.0.1"),
			// 				),
			// 			},
			// IMPORT
			{
				ResourceName:            "dbt_cloud_connection.test_postgres_connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"oauth_client_id", "oauth_client_secret"},
			},
		},
	})
}

func TestAccDbtCloudDatabricksConnectionResource(t *testing.T) {

	connectionName := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	connectionName2 := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	projectName := strings.ToUpper(acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDbtCloudConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDbtCloudConnectionResourcePostgresConfig(connectionName, projectName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbtCloudConnectionExists("dbt_cloud_connection.test_databricks_connection"),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_databricks_connection", "name", connectionName),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_databricks_connection", "host", "test_databricks"),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_databricks_connection", "http_path", "/my/databricks"),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_databricks_connection", "catalog", "moo"),
				),
			},
			// RENAME
			{
				Config: testAccDbtCloudConnectionResourcePostgresConfig(connectionName2, projectName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbtCloudConnectionExists("dbt_cloud_connection.test_databricks_connection"),
					resource.TestCheckResourceAttr("dbt_cloud_connection.test_databricks_connection", "name", connectionName2),
				),
			},
			// 			// MODIFY
			// 			{
			// 				Config: testAccDbtCloudEnvironmentResourceModifiedConfig(projectName, projectName2, environmentName2),
			// 				Check: resource.ComposeTestCheckFunc(
			// 					testAccCheckDbtCloudProjectExists("dbt_cloud_environment.test_env"),
			// 					resource.TestCheckResourceAttr("dbt_cloud_environment.test_env", "name", environmentName2),
			// 					resource.TestCheckResourceAttr("dbt_cloud_environment.test_env", "dbt_version", "1.0.1"),
			// 				),
			// 			},
			// IMPORT
			{
				ResourceName:            "dbt_cloud_connection.test_databricks_connection",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccDbtCloudConnectionResourceBasicConfig(connectionName, projectName, oAuthClientID, oAuthClientSecret string) string {
	return fmt.Sprintf(`
resource "dbt_cloud_project" "test_project" {
  name        = "%s"
}

resource "dbt_cloud_connection" "test_connection" {
  name        = "%s"
  type = "snowflake"
  project_id = dbt_cloud_project.test_project.id
  account = "test"
  database = "db"
  warehouse = "wh"
  role = "user"
  allow_sso = false
  allow_keep_alive = false
  oauth_client_id = "%s"
  oauth_client_secret = "%s"
}
`, projectName, connectionName, oAuthClientID, oAuthClientSecret)
}

func testAccDbtCloudConnectionResourceRedshiftConfig(connectionName, projectName string) string {
	return fmt.Sprintf(`
resource "dbt_cloud_project" "test_project" {
  name        = "%s"
}

resource "dbt_cloud_connection" "test_redshift_connection" {
  name        = "%s"
  type = "redshift"
  project_id = dbt_cloud_project.test_project.id
  host_name = "test_host_name"
  database = "db"
  port = 5432
  tunnel_enabled = true
}
`, projectName, connectionName)
}

func testAccDbtCloudConnectionResourcePostgresConfig(connectionName, projectName string) string {
	return fmt.Sprintf(`
resource "dbt_cloud_project" "test_project" {
  name        = "%s"
}

resource "dbt_cloud_connection" "test_postgres_connection" {
  name        = "%s"
  type = "postgres"
  project_id = dbt_cloud_project.test_project.id
  host_name = "test_postgres"
  database = "db"
  port = 5432
  tunnel_enabled = true
}
`, projectName, connectionName)
}

func testAccDbtCloudConnectionResourceDatabricksConfig(connectionName, projectName string) string {
	return fmt.Sprintf(`
resource "dbt_cloud_project" "test_project" {
  name        = "%s"
}

resource "dbt_cloud_connection" "test_databricks_connection" {
  name        = "%s"
  type = "adapter"
  project_id = dbt_cloud_project.test_project.id
  host_name = "test_databricks"
  http_path = "/my/databricks"
  catalog   = "moo"
}
`, projectName, connectionName)
}

//
// func testAccDbtCloudEnvironmentResourceModifiedConfig(projectName, projectName2, environmentName string) string {
// 	return fmt.Sprintf(`
// resource "dbt_cloud_project" "test_project" {
//   name        = "%s"
// }
//
// resource "dbt_cloud_project" "test_project_2" {
//   name        = "%s"
// }
//
// resource "dbt_cloud_environment" "test_env" {
//   name        = "%s"
//   type = "deployment"
//   dbt_version = "1.0.1"
//   project_id = dbt_cloud_project.test_project_2.id
// }
// `, projectName, projectName2, environmentName)
// }

func testAccCheckDbtCloudConnectionExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		apiClient := testAccProvider.Meta().(*dbt_cloud.Client)
		projectId := strings.Split(rs.Primary.ID, dbt_cloud.ID_DELIMITER)[0]
		connectionId := strings.Split(rs.Primary.ID, dbt_cloud.ID_DELIMITER)[1]

		_, err := apiClient.GetConnection(connectionId, projectId)
		if err != nil {
			return fmt.Errorf("error fetching item with resource %s. %s", resource, err)
		}
		return nil
	}
}

func testAccCheckDbtCloudConnectionDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*dbt_cloud.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dbt_cloud_connection" {
			continue
		}
		projectId := strings.Split(rs.Primary.ID, dbt_cloud.ID_DELIMITER)[0]
		connectionId := strings.Split(rs.Primary.ID, dbt_cloud.ID_DELIMITER)[1]

		_, err := apiClient.GetConnection(connectionId, projectId)
		if err == nil {
			return fmt.Errorf("Connection still exists")
		}
		notFoundErr := "not found"
		expectedErr := regexp.MustCompile(notFoundErr)
		if !expectedErr.Match([]byte(err.Error())) {
			return fmt.Errorf("expected %s, got %s", notFoundErr, err)
		}
	}

	return nil
}
