package workspace_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
)

func TestWorkspace(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Workspace Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = BeforeSuite(func() {
	// Choose a WebDriver:

	if _, ok := os.LookupEnv("HEADLESS"); ok {
		agoutiDriver = agouti.ChromeDriver(
			agouti.ChromeOptions("args", []string{
				// To disable image download during chrome execution.
				// This setting seems to be disable image loading.
				// see:
				//   - https://groups.google.com/a/chromium.org/g/headless-dev/c/0zD4nAyVoCY/m/eZhSgzP2EAAJ
				//   - https://www.chromium.org/blink
				//   - https://chromium.googlesource.com/chromium/blink/+/master/Source/core/frame/Settings.in#256
				//   - https://peter.sh/experiments/chromium-command-line-switches/#blink-settings
				"--blink-settings=imagesEnabled=false",
				"--headless",   // see: https://developers.google.com/web/updates/2017/04/headless-chrome#cli
				"--no-sandbox", // see: https://github.com/theintern/intern/issues/878
			}),
		)
	} else {
		// agoutiDriver = agouti.PhantomJS()
		// agoutiDriver = agouti.Selenium()
		agoutiDriver = agouti.ChromeDriver()
	}
	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
