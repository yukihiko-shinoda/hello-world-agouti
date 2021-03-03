package workspace_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("UserLogin", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("should manage user authentication", func() {
		By("going to the google home page", func() {
			Expect(page.Navigate("https://www.google.com/")).To(Succeed())
			Eventually(page.FindByXPath("//input[@name='q']")).Should(BeFound())
			Expect(page.FindByName("q").Fill("agouti")).To(Succeed())
			enter := "\uE007"
			Expect(page.FindByName("q").SendKeys(enter)).To(Succeed())
			Eventually(page.FindByID("result-stats")).Should(BeVisible())
			Eventually(page.FindByXPath("//a[text()='Red-rumped agouti' and contains(following-sibling::text(), '57\u00A0cm')]")).Should(BeVisible())
			Eventually(page.FindByXPath("//a[text()='Central American agouti' and contains(following-sibling::text(), '2.2\u00A0kg')]")).Should(BeVisible())
			Eventually(page.FindByXPath("//a[text()='Red-rumped agouti' and contains(following-sibling::text(), '3\u00A0kg')]")).Should(BeVisible())
		})
	})
})
