package e2e_test

import (
	"io"
	"net/http"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var cmd *exec.Cmd

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "End-to-End Suite")
}

var _ = BeforeSuite(func() {
	// Start your API server as a subprocess
	cmd = exec.Command("go", "run", "../../main.go")
	err := cmd.Start()
	Expect(err).NotTo(HaveOccurred())

	// Allow time for the server to start
	time.Sleep(1 * time.Second)
})

var _ = AfterSuite(func() {
	// Terminate the subprocess when tests are done
	err := cmd.Process.Kill()
	Expect(err).NotTo(HaveOccurred())
})

var _ = Describe("API Endpoints", func() {
	It("should return 'Hello World!' from /hello", func() {
		resp, err := http.Get("http://localhost:8080/hello")
		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))

		resBody, err := io.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())

		Expect(resBody).Should(Equal([]byte("Hello World!")))
	})
})
