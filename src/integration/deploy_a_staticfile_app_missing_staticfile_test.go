package integration_test

import (
	"integration/cutlass"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("a staticfile app with no staticfile", func() {
	var app *cutlass.App
	AfterEach(func() {
		if app != nil {
			app.Destroy()
		}
		app = nil
	})

	BeforeEach(func() {
		app = cutlass.New(filepath.Join(bpDir, "cf_spec", "fixtures", "without_staticfile"))
		app.Buildpack = "staticfile_buildpack"
	})

	It("runs", func() {
		Expect(app.Push()).To(Succeed())
		Expect(app.InstanceStates()).To(Equal([]string{"RUNNING"}))

		Expect(app.Stdout.String()).ToNot(ContainSubstring("grep: Staticfile: No such file or directory"))
	})
})