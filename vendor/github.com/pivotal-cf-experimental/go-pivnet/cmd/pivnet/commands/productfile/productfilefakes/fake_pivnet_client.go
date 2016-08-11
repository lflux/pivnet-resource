// This file was generated by counterfeiter
package productfilefakes

import (
	"io"
	"sync"

	pivnet "github.com/pivotal-cf-experimental/go-pivnet"
	"github.com/pivotal-cf-experimental/go-pivnet/cmd/pivnet/commands/productfile"
)

type FakePivnetClient struct {
	ReleaseForProductVersionStub        func(productSlug string, releaseVersion string) (pivnet.Release, error)
	releaseForProductVersionMutex       sync.RWMutex
	releaseForProductVersionArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	releaseForProductVersionReturns struct {
		result1 pivnet.Release
		result2 error
	}
	GetProductFilesStub        func(productSlug string) ([]pivnet.ProductFile, error)
	getProductFilesMutex       sync.RWMutex
	getProductFilesArgsForCall []struct {
		productSlug string
	}
	getProductFilesReturns struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	GetProductFilesForReleaseStub        func(productSlug string, releaseID int) ([]pivnet.ProductFile, error)
	getProductFilesForReleaseMutex       sync.RWMutex
	getProductFilesForReleaseArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	getProductFilesForReleaseReturns struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	GetProductFileStub        func(productSlug string, productFileID int) (pivnet.ProductFile, error)
	getProductFileMutex       sync.RWMutex
	getProductFileArgsForCall []struct {
		productSlug   string
		productFileID int
	}
	getProductFileReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	GetProductFileForReleaseStub        func(productSlug string, releaseID int, productFileID int) (pivnet.ProductFile, error)
	getProductFileForReleaseMutex       sync.RWMutex
	getProductFileForReleaseArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	getProductFileForReleaseReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	AddProductFileStub        func(productSlug string, releaseID int, productFileID int) error
	addProductFileMutex       sync.RWMutex
	addProductFileArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	addProductFileReturns struct {
		result1 error
	}
	RemoveProductFileStub        func(productSlug string, releaseID int, productFileID int) error
	removeProductFileMutex       sync.RWMutex
	removeProductFileArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	removeProductFileReturns struct {
		result1 error
	}
	DeleteProductFileStub        func(productSlug string, releaseID int) (pivnet.ProductFile, error)
	deleteProductFileMutex       sync.RWMutex
	deleteProductFileArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	deleteProductFileReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	AcceptEULAStub        func(productSlug string, releaseID int) error
	acceptEULAMutex       sync.RWMutex
	acceptEULAArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	acceptEULAReturns struct {
		result1 error
	}
	DownloadFileStub        func(writer io.Writer, downloadLink string) error
	downloadFileMutex       sync.RWMutex
	downloadFileArgsForCall []struct {
		writer       io.Writer
		downloadLink string
	}
	downloadFileReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) ReleaseForProductVersion(productSlug string, releaseVersion string) (pivnet.Release, error) {
	fake.releaseForProductVersionMutex.Lock()
	fake.releaseForProductVersionArgsForCall = append(fake.releaseForProductVersionArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("ReleaseForProductVersion", []interface{}{productSlug, releaseVersion})
	fake.releaseForProductVersionMutex.Unlock()
	if fake.ReleaseForProductVersionStub != nil {
		return fake.ReleaseForProductVersionStub(productSlug, releaseVersion)
	} else {
		return fake.releaseForProductVersionReturns.result1, fake.releaseForProductVersionReturns.result2
	}
}

func (fake *FakePivnetClient) ReleaseForProductVersionCallCount() int {
	fake.releaseForProductVersionMutex.RLock()
	defer fake.releaseForProductVersionMutex.RUnlock()
	return len(fake.releaseForProductVersionArgsForCall)
}

func (fake *FakePivnetClient) ReleaseForProductVersionArgsForCall(i int) (string, string) {
	fake.releaseForProductVersionMutex.RLock()
	defer fake.releaseForProductVersionMutex.RUnlock()
	return fake.releaseForProductVersionArgsForCall[i].productSlug, fake.releaseForProductVersionArgsForCall[i].releaseVersion
}

func (fake *FakePivnetClient) ReleaseForProductVersionReturns(result1 pivnet.Release, result2 error) {
	fake.ReleaseForProductVersionStub = nil
	fake.releaseForProductVersionReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) GetProductFiles(productSlug string) ([]pivnet.ProductFile, error) {
	fake.getProductFilesMutex.Lock()
	fake.getProductFilesArgsForCall = append(fake.getProductFilesArgsForCall, struct {
		productSlug string
	}{productSlug})
	fake.recordInvocation("GetProductFiles", []interface{}{productSlug})
	fake.getProductFilesMutex.Unlock()
	if fake.GetProductFilesStub != nil {
		return fake.GetProductFilesStub(productSlug)
	} else {
		return fake.getProductFilesReturns.result1, fake.getProductFilesReturns.result2
	}
}

func (fake *FakePivnetClient) GetProductFilesCallCount() int {
	fake.getProductFilesMutex.RLock()
	defer fake.getProductFilesMutex.RUnlock()
	return len(fake.getProductFilesArgsForCall)
}

func (fake *FakePivnetClient) GetProductFilesArgsForCall(i int) string {
	fake.getProductFilesMutex.RLock()
	defer fake.getProductFilesMutex.RUnlock()
	return fake.getProductFilesArgsForCall[i].productSlug
}

func (fake *FakePivnetClient) GetProductFilesReturns(result1 []pivnet.ProductFile, result2 error) {
	fake.GetProductFilesStub = nil
	fake.getProductFilesReturns = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) GetProductFilesForRelease(productSlug string, releaseID int) ([]pivnet.ProductFile, error) {
	fake.getProductFilesForReleaseMutex.Lock()
	fake.getProductFilesForReleaseArgsForCall = append(fake.getProductFilesForReleaseArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("GetProductFilesForRelease", []interface{}{productSlug, releaseID})
	fake.getProductFilesForReleaseMutex.Unlock()
	if fake.GetProductFilesForReleaseStub != nil {
		return fake.GetProductFilesForReleaseStub(productSlug, releaseID)
	} else {
		return fake.getProductFilesForReleaseReturns.result1, fake.getProductFilesForReleaseReturns.result2
	}
}

func (fake *FakePivnetClient) GetProductFilesForReleaseCallCount() int {
	fake.getProductFilesForReleaseMutex.RLock()
	defer fake.getProductFilesForReleaseMutex.RUnlock()
	return len(fake.getProductFilesForReleaseArgsForCall)
}

func (fake *FakePivnetClient) GetProductFilesForReleaseArgsForCall(i int) (string, int) {
	fake.getProductFilesForReleaseMutex.RLock()
	defer fake.getProductFilesForReleaseMutex.RUnlock()
	return fake.getProductFilesForReleaseArgsForCall[i].productSlug, fake.getProductFilesForReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) GetProductFilesForReleaseReturns(result1 []pivnet.ProductFile, result2 error) {
	fake.GetProductFilesForReleaseStub = nil
	fake.getProductFilesForReleaseReturns = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) GetProductFile(productSlug string, productFileID int) (pivnet.ProductFile, error) {
	fake.getProductFileMutex.Lock()
	fake.getProductFileArgsForCall = append(fake.getProductFileArgsForCall, struct {
		productSlug   string
		productFileID int
	}{productSlug, productFileID})
	fake.recordInvocation("GetProductFile", []interface{}{productSlug, productFileID})
	fake.getProductFileMutex.Unlock()
	if fake.GetProductFileStub != nil {
		return fake.GetProductFileStub(productSlug, productFileID)
	} else {
		return fake.getProductFileReturns.result1, fake.getProductFileReturns.result2
	}
}

func (fake *FakePivnetClient) GetProductFileCallCount() int {
	fake.getProductFileMutex.RLock()
	defer fake.getProductFileMutex.RUnlock()
	return len(fake.getProductFileArgsForCall)
}

func (fake *FakePivnetClient) GetProductFileArgsForCall(i int) (string, int) {
	fake.getProductFileMutex.RLock()
	defer fake.getProductFileMutex.RUnlock()
	return fake.getProductFileArgsForCall[i].productSlug, fake.getProductFileArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) GetProductFileReturns(result1 pivnet.ProductFile, result2 error) {
	fake.GetProductFileStub = nil
	fake.getProductFileReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) GetProductFileForRelease(productSlug string, releaseID int, productFileID int) (pivnet.ProductFile, error) {
	fake.getProductFileForReleaseMutex.Lock()
	fake.getProductFileForReleaseArgsForCall = append(fake.getProductFileForReleaseArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("GetProductFileForRelease", []interface{}{productSlug, releaseID, productFileID})
	fake.getProductFileForReleaseMutex.Unlock()
	if fake.GetProductFileForReleaseStub != nil {
		return fake.GetProductFileForReleaseStub(productSlug, releaseID, productFileID)
	} else {
		return fake.getProductFileForReleaseReturns.result1, fake.getProductFileForReleaseReturns.result2
	}
}

func (fake *FakePivnetClient) GetProductFileForReleaseCallCount() int {
	fake.getProductFileForReleaseMutex.RLock()
	defer fake.getProductFileForReleaseMutex.RUnlock()
	return len(fake.getProductFileForReleaseArgsForCall)
}

func (fake *FakePivnetClient) GetProductFileForReleaseArgsForCall(i int) (string, int, int) {
	fake.getProductFileForReleaseMutex.RLock()
	defer fake.getProductFileForReleaseMutex.RUnlock()
	return fake.getProductFileForReleaseArgsForCall[i].productSlug, fake.getProductFileForReleaseArgsForCall[i].releaseID, fake.getProductFileForReleaseArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) GetProductFileForReleaseReturns(result1 pivnet.ProductFile, result2 error) {
	fake.GetProductFileForReleaseStub = nil
	fake.getProductFileForReleaseReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AddProductFile(productSlug string, releaseID int, productFileID int) error {
	fake.addProductFileMutex.Lock()
	fake.addProductFileArgsForCall = append(fake.addProductFileArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("AddProductFile", []interface{}{productSlug, releaseID, productFileID})
	fake.addProductFileMutex.Unlock()
	if fake.AddProductFileStub != nil {
		return fake.AddProductFileStub(productSlug, releaseID, productFileID)
	} else {
		return fake.addProductFileReturns.result1
	}
}

func (fake *FakePivnetClient) AddProductFileCallCount() int {
	fake.addProductFileMutex.RLock()
	defer fake.addProductFileMutex.RUnlock()
	return len(fake.addProductFileArgsForCall)
}

func (fake *FakePivnetClient) AddProductFileArgsForCall(i int) (string, int, int) {
	fake.addProductFileMutex.RLock()
	defer fake.addProductFileMutex.RUnlock()
	return fake.addProductFileArgsForCall[i].productSlug, fake.addProductFileArgsForCall[i].releaseID, fake.addProductFileArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) AddProductFileReturns(result1 error) {
	fake.AddProductFileStub = nil
	fake.addProductFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveProductFile(productSlug string, releaseID int, productFileID int) error {
	fake.removeProductFileMutex.Lock()
	fake.removeProductFileArgsForCall = append(fake.removeProductFileArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("RemoveProductFile", []interface{}{productSlug, releaseID, productFileID})
	fake.removeProductFileMutex.Unlock()
	if fake.RemoveProductFileStub != nil {
		return fake.RemoveProductFileStub(productSlug, releaseID, productFileID)
	} else {
		return fake.removeProductFileReturns.result1
	}
}

func (fake *FakePivnetClient) RemoveProductFileCallCount() int {
	fake.removeProductFileMutex.RLock()
	defer fake.removeProductFileMutex.RUnlock()
	return len(fake.removeProductFileArgsForCall)
}

func (fake *FakePivnetClient) RemoveProductFileArgsForCall(i int) (string, int, int) {
	fake.removeProductFileMutex.RLock()
	defer fake.removeProductFileMutex.RUnlock()
	return fake.removeProductFileArgsForCall[i].productSlug, fake.removeProductFileArgsForCall[i].releaseID, fake.removeProductFileArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) RemoveProductFileReturns(result1 error) {
	fake.RemoveProductFileStub = nil
	fake.removeProductFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) DeleteProductFile(productSlug string, releaseID int) (pivnet.ProductFile, error) {
	fake.deleteProductFileMutex.Lock()
	fake.deleteProductFileArgsForCall = append(fake.deleteProductFileArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("DeleteProductFile", []interface{}{productSlug, releaseID})
	fake.deleteProductFileMutex.Unlock()
	if fake.DeleteProductFileStub != nil {
		return fake.DeleteProductFileStub(productSlug, releaseID)
	} else {
		return fake.deleteProductFileReturns.result1, fake.deleteProductFileReturns.result2
	}
}

func (fake *FakePivnetClient) DeleteProductFileCallCount() int {
	fake.deleteProductFileMutex.RLock()
	defer fake.deleteProductFileMutex.RUnlock()
	return len(fake.deleteProductFileArgsForCall)
}

func (fake *FakePivnetClient) DeleteProductFileArgsForCall(i int) (string, int) {
	fake.deleteProductFileMutex.RLock()
	defer fake.deleteProductFileMutex.RUnlock()
	return fake.deleteProductFileArgsForCall[i].productSlug, fake.deleteProductFileArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) DeleteProductFileReturns(result1 pivnet.ProductFile, result2 error) {
	fake.DeleteProductFileStub = nil
	fake.deleteProductFileReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AcceptEULA(productSlug string, releaseID int) error {
	fake.acceptEULAMutex.Lock()
	fake.acceptEULAArgsForCall = append(fake.acceptEULAArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("AcceptEULA", []interface{}{productSlug, releaseID})
	fake.acceptEULAMutex.Unlock()
	if fake.AcceptEULAStub != nil {
		return fake.AcceptEULAStub(productSlug, releaseID)
	} else {
		return fake.acceptEULAReturns.result1
	}
}

func (fake *FakePivnetClient) AcceptEULACallCount() int {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return len(fake.acceptEULAArgsForCall)
}

func (fake *FakePivnetClient) AcceptEULAArgsForCall(i int) (string, int) {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return fake.acceptEULAArgsForCall[i].productSlug, fake.acceptEULAArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) AcceptEULAReturns(result1 error) {
	fake.AcceptEULAStub = nil
	fake.acceptEULAReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) DownloadFile(writer io.Writer, downloadLink string) error {
	fake.downloadFileMutex.Lock()
	fake.downloadFileArgsForCall = append(fake.downloadFileArgsForCall, struct {
		writer       io.Writer
		downloadLink string
	}{writer, downloadLink})
	fake.recordInvocation("DownloadFile", []interface{}{writer, downloadLink})
	fake.downloadFileMutex.Unlock()
	if fake.DownloadFileStub != nil {
		return fake.DownloadFileStub(writer, downloadLink)
	} else {
		return fake.downloadFileReturns.result1
	}
}

func (fake *FakePivnetClient) DownloadFileCallCount() int {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return len(fake.downloadFileArgsForCall)
}

func (fake *FakePivnetClient) DownloadFileArgsForCall(i int) (io.Writer, string) {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return fake.downloadFileArgsForCall[i].writer, fake.downloadFileArgsForCall[i].downloadLink
}

func (fake *FakePivnetClient) DownloadFileReturns(result1 error) {
	fake.DownloadFileStub = nil
	fake.downloadFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releaseForProductVersionMutex.RLock()
	defer fake.releaseForProductVersionMutex.RUnlock()
	fake.getProductFilesMutex.RLock()
	defer fake.getProductFilesMutex.RUnlock()
	fake.getProductFilesForReleaseMutex.RLock()
	defer fake.getProductFilesForReleaseMutex.RUnlock()
	fake.getProductFileMutex.RLock()
	defer fake.getProductFileMutex.RUnlock()
	fake.getProductFileForReleaseMutex.RLock()
	defer fake.getProductFileForReleaseMutex.RUnlock()
	fake.addProductFileMutex.RLock()
	defer fake.addProductFileMutex.RUnlock()
	fake.removeProductFileMutex.RLock()
	defer fake.removeProductFileMutex.RUnlock()
	fake.deleteProductFileMutex.RLock()
	defer fake.deleteProductFileMutex.RUnlock()
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ productfile.PivnetClient = new(FakePivnetClient)