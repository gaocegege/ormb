package model

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Format", func() {
	It("Should validate SavedModel format successfully", func() {
		savedmodelFormat := FormatSavedModel
		err := savedmodelFormat.ValidateDirectory("../../examples/SavedModel-fashion")
		Expect(err).To(BeNil())
	})

	It("Should validate ONNX format successfully", func() {
		savedmodelFormat := FormatONNX
		err := savedmodelFormat.ValidateDirectory("../../examples/ONNX-model")
		Expect(err).To(BeNil())
	})

	It("Should validate H5 format successfully", func() {
		savedmodelFormat := FormatH5
		err := savedmodelFormat.ValidateDirectory("../../examples/H5-model")
		Expect(err).To(BeNil())
	})

	It("Should validate PMML format successfully", func() {
		savedmodelFormat := FormatPMML
		err := savedmodelFormat.ValidateDirectory("../../examples/PMML-model")
		Expect(err).To(BeNil())
	})

	It("Should validate CaffeModel format successfully", func() {
		savedmodelFormat := FormatCaffeModel
		err := savedmodelFormat.ValidateDirectory("../../examples/Caffe-model")
		Expect(err).To(BeNil())
	})

	It("Should validate NetDef format successfully", func() {
		savedmodelFormat := FormatNetDef
		err := savedmodelFormat.ValidateDirectory("../../examples/Caffe2-model")
		Expect(err).To(BeNil())
	})

	It("Should validate MXNETParams format successfully", func() {
		savedmodelFormat := FormatMXNETParams
		err := savedmodelFormat.ValidateDirectory("../../examples/MXNETParams-model")
		Expect(err).To(BeNil())
	})

	It("Should validate TorchScript format successfully", func() {
		savedmodelFormat := FormatTorchScript
		err := savedmodelFormat.ValidateDirectory("../../examples/TorchScript-model")
		Expect(err).To(BeNil())
	})

	It("Should validate GraphDef format successfully", func() {
		savedmodelFormat := FormatGraphDef
		err := savedmodelFormat.ValidateDirectory("../../examples/GraphDef-model")
		Expect(err).To(BeNil())
	})
})
