package redis_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccRedisInstanceDatasource_basic(t *testing.T) {
	t.Parallel()

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRedisInstanceDatasourceConfig(acctest.RandString(t, 10)),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckDataSourceStateMatchesResourceState("data.google_redis_instance.redis", "google_redis_instance.redis"),
					resource.TestCheckResourceAttrSet("data.google_redis_instance.redis", "reserved_ip_range"),
				),
			},
		},
	})
}

func testAccRedisInstanceDatasourceConfig(suffix string) string {
	return fmt.Sprintf(`
resource "google_redis_instance" "redis" {
  name               = "redis-test-%s"
  memory_size_gb     = 1

  labels   = {
    my-label = "my-label-value"
  }
}

data "google_redis_instance" "redis" {
  name = google_redis_instance.redis.name
}
`, suffix)
}
