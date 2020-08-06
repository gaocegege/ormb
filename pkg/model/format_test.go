package model_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/kleveross/ormb/pkg/model"
)

var _ = Describe("Format", func() {
	It("Should validate SavedModel format successfully", func() {
		savedmodelFormat := model.FormatSavedModel
		err := savedmodelFormat.ValidateDirectory("../../examples/SavedModel-fashion")
		Expect(err).To(BeNil())
	})

	It("Should validate ONNX format successfully", func() {
		savedmodelFormat := model.FormatONNX
		err := savedmodelFormat.ValidateDirectory("../../examples/ONNX-model")
		Expect(err).To(BeNil())
	})

	It("Should validate H5 format successfully", func() {
		savedmodelFormat := model.FormatH5
		err := savedmodelFormat.ValidateDirectory("../../examples/H5-model")
		Expect(err).To(BeNil())
	})

	It("Should validate PMML format successfully", func() {
		savedmodelFormat := model.FormatPMML
		err := savedmodelFormat.ValidateDirectory("../../examples/PMML-model")
		Expect(err).To(BeNil())
	})

	It("Should validate CaffeModel format successfully", func() {
		savedmodelFormat := model.FormatCaffeModel
		err := savedmodelFormat.ValidateDirectory("../../examples/Caffe-model")
		Expect(err).To(BeNil())
	})

	It("Should validate NetDef format successfully", func() {
		savedmodelFormat := model.FormatNetDef
		err := savedmodelFormat.ValidateDirectory("../../examples/Caffe2-model")
		Expect(err).To(BeNil())
	})

	It("Should validate MXNETParams format successfully", func() {
		savedmodelFormat := model.FormatMXNETParams
		err := savedmodelFormat.ValidateDirectory("../../examples/MXNETParams-model")
		Expect(err).To(BeNil())
	})

	It("Should validate TorchScript format successfully", func() {
		savedmodelFormat := model.FormatTorchScript
		err := savedmodelFormat.ValidateDirectory("../../examples/TorchScript-model")
		Expect(err).To(BeNil())
	})

	It("Should validate GraphDef format successfully", func() {
		savedmodelFormat := model.FormatGraphDef
		err := savedmodelFormat.ValidateDirectory("../../examples/GraphDef-model")
		Expect(err).To(BeNil())
	})
})
