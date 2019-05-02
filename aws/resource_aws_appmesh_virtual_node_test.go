package aws

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func testAccAwsAppmeshVirtualNode_basic(t *testing.T) {
	var vn appmesh.VirtualNodeData
	resourceName := "aws_appmesh_virtual_node.foo"
	meshName := fmt.Sprintf("tf-test-mesh-%d", acctest.RandInt())
	vnName := fmt.Sprintf("tf-test-node-%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppmeshVirtualNodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppmeshVirtualNodeConfig_basic(meshName, vnName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppmeshVirtualNodeExists(resourceName, &vn),
					resource.TestCheckResourceAttr(resourceName, "name", vnName),
					resource.TestCheckResourceAttr(resourceName, "mesh_name", meshName),
					resource.TestCheckResourceAttr(resourceName, "spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.backend.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.logging.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.service_discovery.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "created_date"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated_date"),
					resource.TestMatchResourceAttr(resourceName, "arn", regexp.MustCompile(fmt.Sprintf("^arn:[^:]+:appmesh:[^:]+:\\d{12}:mesh/%s/virtualNode/%s", meshName, vnName))),
				),
			},
			{
				ResourceName:      resourceName,
				ImportStateId:     fmt.Sprintf("%s/%s", meshName, vnName),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAwsAppmeshVirtualNode_listenerHealthChecks(t *testing.T) {
	var vn appmesh.VirtualNodeData
	resourceName := "aws_appmesh_virtual_node.foo"
	meshName := fmt.Sprintf("tf-test-mesh-%d", acctest.RandInt())
	vnName := fmt.Sprintf("tf-test-node-%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppmeshVirtualNodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppmeshVirtualNodeConfig_listenerHealthChecks(meshName, vnName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppmeshVirtualNodeExists(resourceName, &vn),
					resource.TestCheckResourceAttr(resourceName, "name", vnName),
					resource.TestCheckResourceAttr(resourceName, "mesh_name", meshName),
					resource.TestCheckResourceAttr(resourceName, "spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.backend.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.backend.2622272660.virtual_service.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.backend.2622272660.virtual_service.0.virtual_service_name", "servicea.simpleapp.local"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.433446196.health_check.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.433446196.health_check.0.healthy_threshold", "3"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.433446196.health_check.0.interval_millis", "5000"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.433446196.health_check.0.path", "/ping"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.433446196.health_check.0.port", "8080"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.433446196.health_check.0.protocol", "http"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.433446196.health_check.0.timeout_millis", "2000"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.433446196.health_check.0.unhealthy_threshold", "5"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.433446196.port_mapping.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.433446196.port_mapping.0.port", "8080"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.433446196.port_mapping.0.protocol", "http"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.logging.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.service_discovery.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.service_discovery.0.dns.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.service_discovery.0.dns.0.hostname", "serviceb.simpleapp.local"),
					resource.TestCheckResourceAttrSet(resourceName, "created_date"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated_date"),
					resource.TestMatchResourceAttr(resourceName, "arn", regexp.MustCompile(fmt.Sprintf("^arn:[^:]+:appmesh:[^:]+:\\d{12}:mesh/%s/virtualNode/%s", meshName, vnName))),
				),
			},
			{
				Config: testAccAppmeshVirtualNodeConfig_listenerHealthChecksUpdated(meshName, vnName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppmeshVirtualNodeExists(resourceName, &vn),
					resource.TestCheckResourceAttr(resourceName, "name", vnName),
					resource.TestCheckResourceAttr(resourceName, "mesh_name", meshName),
					resource.TestCheckResourceAttr(resourceName, "spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.backend.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.backend.2576932631.virtual_service.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.backend.2576932631.virtual_service.0.virtual_service_name", "servicec.simpleapp.local"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.backend.2025248115.virtual_service.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.backend.2025248115.virtual_service.0.virtual_service_name", "serviced.simpleapp.local"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.3446683576.health_check.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.3446683576.health_check.0.healthy_threshold", "4"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.3446683576.health_check.0.interval_millis", "7000"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.3446683576.health_check.0.port", "8081"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.3446683576.health_check.0.protocol", "tcp"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.3446683576.health_check.0.timeout_millis", "3000"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.3446683576.health_check.0.unhealthy_threshold", "9"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.3446683576.port_mapping.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.3446683576.port_mapping.0.port", "8081"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.listener.3446683576.port_mapping.0.protocol", "http"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.logging.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.service_discovery.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.service_discovery.0.dns.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.service_discovery.0.dns.0.hostname", "serviceb1.simpleapp.local"),
					resource.TestCheckResourceAttrSet(resourceName, "created_date"),
					resource.TestCheckResourceAttrSet(resourceName, "last_updated_date"),
					resource.TestMatchResourceAttr(resourceName, "arn", regexp.MustCompile(fmt.Sprintf("^arn:[^:]+:appmesh:[^:]+:\\d{12}:mesh/%s/virtualNode/%s", meshName, vnName))),
				),
			},
			{
				ResourceName:      resourceName,
				ImportStateId:     fmt.Sprintf("%s/%s", meshName, vnName),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAwsAppmeshVirtualNode_logging(t *testing.T) {
	var vn appmesh.VirtualNodeData
	resourceName := "aws_appmesh_virtual_node.foo"
	meshName := fmt.Sprintf("tf-test-mesh-%d", acctest.RandInt())
	vnName := fmt.Sprintf("tf-test-node-%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppmeshVirtualNodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppmeshVirtualNodeConfig_logging(meshName, vnName, "/dev/stdout"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppmeshVirtualNodeExists(resourceName, &vn),
					resource.TestCheckResourceAttr(resourceName, "name", vnName),
					resource.TestCheckResourceAttr(resourceName, "mesh_name", meshName),
					resource.TestCheckResourceAttr(resourceName, "spec.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.logging.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.logging.0.access_log.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.logging.0.access_log.0.file.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.logging.0.access_log.0.file.0.path", "/dev/stdout"),
				),
			},
			{
				Config: testAccAppmeshVirtualNodeConfig_logging(meshName, vnName, "/tmp/access.log"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppmeshVirtualNodeExists(resourceName, &vn),
					resource.TestCheckResourceAttr(resourceName, "name", vnName),
					resource.TestCheckResourceAttr(resourceName, "mesh_name", meshName),
					resource.TestCheckResourceAttr(resourceName, "spec.0.logging.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.logging.0.access_log.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.logging.0.access_log.0.file.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "spec.0.logging.0.access_log.0.file.0.path", "/tmp/access.log"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportStateId:     fmt.Sprintf("%s/%s", meshName, vnName),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckAppmeshVirtualNodeDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).appmeshconn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_appmesh_virtual_node" {
			continue
		}

		_, err := conn.DescribeVirtualNode(&appmesh.DescribeVirtualNodeInput{
			MeshName:        aws.String(rs.Primary.Attributes["mesh_name"]),
			VirtualNodeName: aws.String(rs.Primary.Attributes["name"]),
		})
		if isAWSErr(err, appmesh.ErrCodeNotFoundException, "") {
			continue
		}
		if err != nil {
			return err
		}
		return fmt.Errorf("still exist.")
	}

	return nil
}

func testAccCheckAppmeshVirtualNodeExists(name string, v *appmesh.VirtualNodeData) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*AWSClient).appmeshconn

		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		resp, err := conn.DescribeVirtualNode(&appmesh.DescribeVirtualNodeInput{
			MeshName:        aws.String(rs.Primary.Attributes["mesh_name"]),
			VirtualNodeName: aws.String(rs.Primary.Attributes["name"]),
		})
		if err != nil {
			return err
		}

		*v = *resp.VirtualNode

		return nil
	}
}

func testAccAppmeshVirtualNodeConfig_basic(meshName, vnName string) string {
	return fmt.Sprintf(`
resource "aws_appmesh_mesh" "foo" {
  name = %[1]q
}

resource "aws_appmesh_virtual_node" "foo" {
  name      = %[2]q
  mesh_name = "${aws_appmesh_mesh.foo.id}"

  spec {}
}
`, meshName, vnName)
}

func testAccAppmeshVirtualNodeConfig_listenerHealthChecks(meshName, vnName string) string {
	return fmt.Sprintf(`
resource "aws_appmesh_mesh" "foo" {
  name = %[1]q
}

resource "aws_appmesh_virtual_node" "foo" {
  name      = %[2]q
  mesh_name = "${aws_appmesh_mesh.foo.id}"

  spec {
    backend {
      virtual_service {
        virtual_service_name = "servicea.simpleapp.local"
      }
    }

    listener {
      port_mapping {
        port     = 8080
        protocol = "http"
      }

      health_check {
        protocol            = "http"
        path                = "/ping"
        healthy_threshold   = 3
        unhealthy_threshold = 5
        timeout_millis      = 2000
        interval_millis     = 5000
      }
    }

    service_discovery {
      dns {
        hostname = "serviceb.simpleapp.local"
      }
    }
  }
}
`, meshName, vnName)
}

func testAccAppmeshVirtualNodeConfig_listenerHealthChecksUpdated(meshName, vnName string) string {
	return fmt.Sprintf(`
resource "aws_appmesh_mesh" "foo" {
  name = %[1]q
}

resource "aws_appmesh_virtual_node" "foo" {
  name      = %[2]q
  mesh_name = "${aws_appmesh_mesh.foo.id}"

  spec {
    backend {
      virtual_service {
        virtual_service_name = "servicec.simpleapp.local"
      }
    }

    backend {
      virtual_service {
        virtual_service_name = "serviced.simpleapp.local"
      }
    }

    listener {
      port_mapping {
        port     = 8081
        protocol = "http"
      }

      health_check {
        protocol            = "tcp"
        port                = 8081
        healthy_threshold   = 4
        unhealthy_threshold = 9
        timeout_millis      = 3000
        interval_millis     = 7000
      }
    }

    service_discovery {
      dns {
        hostname = "serviceb1.simpleapp.local"
      }
    }
  }
}
`, meshName, vnName)
}

func testAccAppmeshVirtualNodeConfig_logging(meshName, vnName, path string) string {
	return fmt.Sprintf(`
resource "aws_appmesh_mesh" "foo" {
  name = %[1]q
}

resource "aws_appmesh_virtual_node" "foo" {
  name      = %[2]q
  mesh_name = "${aws_appmesh_mesh.foo.id}"

  spec {
    backend {
      virtual_service {
        virtual_service_name = "servicea.simpleapp.local"
      }
    }

    listener {
      port_mapping {
        port     = 8080
        protocol = "http"
      }
    }

    logging {
      access_log {
        file {
          path = %[3]q
        }
      }
    }

    service_discovery {
      dns {
        hostname = "serviceb.simpleapp.local"
      }
    }
  }
}
`, meshName, vnName, path)
}