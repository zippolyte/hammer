package integration

import (
	"io/ioutil"
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("CF", func() {
	It("generates the correct script", func() {
		command := exec.Command(pathToPcf, "cf-login", "-t", "fixtures/claim_manatee_response.json", "-f")
		session, err := Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session).Should(Exit(0))
		Eventually(string(session.Err.Contents())).Should(Equal(""))
		Eventually(session.Out).Should(Say("Logging in to: https://pcf.manatee.cf-app.com"))

		output := strings.TrimSuffix(string(session.Out.Contents()), "\n")
		lines := strings.Split(output, "\n")
		pathToFile := lines[len(lines)-1]
		contents, err := ioutil.ReadFile(pathToFile)
		Expect(err).NotTo(HaveOccurred())

		Expect(string(contents)).To(Equal(LoadFixture("cf_script.sh")))
	})
})
