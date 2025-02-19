package sandbox_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	types "k8s.io/cri-api/pkg/apis/runtime/v1"
)

var _ = Describe("Sandbox", func() {
	t.Describe("SetConfig", func() {
		It("should succeed", func() {
			// Given
			config := &types.PodSandboxConfig{
				Metadata: &types.PodSandboxMetadata{Name: "name"},
			}

			// When
			err := sut.SetConfig(config)

			// Then
			Expect(err).To(BeNil())
			Expect(sut.Config()).To(Equal(config))
		})

		It("should fail with nil config", func() {
			// Given
			// When
			err := sut.SetConfig(nil)

			// Then
			Expect(err).NotTo(BeNil())
			Expect(sut.Config()).To(BeNil())
		})

		It("should fail with empty config", func() {
			// Given
			config := &types.PodSandboxConfig{}

			// When
			err := sut.SetConfig(config)

			// Then
			Expect(err).NotTo(BeNil())
			Expect(sut.Config()).To(BeNil())
		})

		It("should fail with an empty name", func() {
			// Given
			config := &types.PodSandboxConfig{
				Metadata: &types.PodSandboxMetadata{},
			}

			// When
			err := sut.SetConfig(config)

			// Then
			Expect(err).NotTo(BeNil())
			Expect(sut.Config()).To(BeNil())
		})

		It("should fail with config already set", func() {
			// Given
			config := &types.PodSandboxConfig{
				Metadata: &types.PodSandboxMetadata{Name: "name"},
			}
			err := sut.SetConfig(config)
			Expect(err).To(BeNil())

			// When
			err = sut.SetConfig(config)

			// Then
			Expect(err).NotTo(BeNil())
			Expect(sut.Config()).NotTo(BeNil())
		})
	})
})
