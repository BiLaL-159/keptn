// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package common_mock

import (
	"context"
	"github.com/go-git/go-git/v5"
	config2 "github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"sync"
)

// RepositoryMock is a mock implementation of common.Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked common.Repository
// 		mockedRepository := &RepositoryMock{
// 			BlobObjectFunc: func(h plumbing.Hash) (*object.Blob, error) {
// 				panic("mock out the BlobObject method")
// 			},
// 			BlobObjectsFunc: func() (*object.BlobIter, error) {
// 				panic("mock out the BlobObjects method")
// 			},
// 			BranchFunc: func(name string) (*config2.Branch, error) {
// 				panic("mock out the Branch method")
// 			},
// 			BranchesFunc: func() (storer.ReferenceIter, error) {
// 				panic("mock out the Branches method")
// 			},
// 			CommitObjectFunc: func(h plumbing.Hash) (*object.Commit, error) {
// 				panic("mock out the CommitObject method")
// 			},
// 			CommitObjectsFunc: func() (object.CommitIter, error) {
// 				panic("mock out the CommitObjects method")
// 			},
// 			CreateBranchFunc: func(c *config2.Branch) error {
// 				panic("mock out the CreateBranch method")
// 			},
// 			CreateRemoteFunc: func(c *config2.RemoteConfig) (*git.Remote, error) {
// 				panic("mock out the CreateRemote method")
// 			},
// 			CreateRemoteAnonymousFunc: func(c *config2.RemoteConfig) (*git.Remote, error) {
// 				panic("mock out the CreateRemoteAnonymous method")
// 			},
// 			DeleteBranchFunc: func(name string) error {
// 				panic("mock out the DeleteBranch method")
// 			},
// 			DeleteRemoteFunc: func(name string) error {
// 				panic("mock out the DeleteRemote method")
// 			},
// 			HeadFunc: func() (*plumbing.Reference, error) {
// 				panic("mock out the Head method")
// 			},
// 			ObjectFunc: func(t plumbing.ObjectType, h plumbing.Hash) (object.Object, error) {
// 				panic("mock out the Object method")
// 			},
// 			ObjectsFunc: func() (*object.ObjectIter, error) {
// 				panic("mock out the Objects method")
// 			},
// 			PushFunc: func(o *git.PushOptions) error {
// 				panic("mock out the Push method")
// 			},
// 			ReferenceFunc: func(name plumbing.ReferenceName, resolved bool) (*plumbing.Reference, error) {
// 				panic("mock out the Reference method")
// 			},
// 			ReferencesFunc: func() (storer.ReferenceIter, error) {
// 				panic("mock out the References method")
// 			},
// 			RemoteFunc: func(name string) (*git.Remote, error) {
// 				panic("mock out the Remote method")
// 			},
// 			RemotesFunc: func() ([]*git.Remote, error) {
// 				panic("mock out the Remotes method")
// 			},
// 			ResolveRevisionFunc: func(rev plumbing.Revision) (*plumbing.Hash, error) {
// 				panic("mock out the ResolveRevision method")
// 			},
// 			TreeObjectFunc: func(h plumbing.Hash) (*object.Tree, error) {
// 				panic("mock out the TreeObject method")
// 			},
// 			TreeObjectsFunc: func() (*object.TreeIter, error) {
// 				panic("mock out the TreeObjects method")
// 			},
// 			WorktreeFunc: func() (*git.Worktree, error) {
// 				panic("mock out the Worktree method")
// 			},
// 			calculateRemoteHeadReferenceFunc: func(spec []config2.RefSpec, resolvedHead *plumbing.Reference) []*plumbing.Reference {
// 				panic("mock out the calculateRemoteHeadReference method")
// 			},
// 			cloneFunc: func(ctx context.Context, o *git.CloneOptions) error {
// 				panic("mock out the clone method")
// 			},
// 			resolveToCommitHashFunc: func(h plumbing.Hash) (plumbing.Hash, error) {
// 				panic("mock out the resolveToCommitHash method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires common.Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// BlobObjectFunc mocks the BlobObject method.
	BlobObjectFunc func(h plumbing.Hash) (*object.Blob, error)

	// BlobObjectsFunc mocks the BlobObjects method.
	BlobObjectsFunc func() (*object.BlobIter, error)

	// BranchFunc mocks the Branch method.
	BranchFunc func(name string) (*config2.Branch, error)

	// BranchesFunc mocks the Branches method.
	BranchesFunc func() (storer.ReferenceIter, error)

	// CommitObjectFunc mocks the CommitObject method.
	CommitObjectFunc func(h plumbing.Hash) (*object.Commit, error)

	// CommitObjectsFunc mocks the CommitObjects method.
	CommitObjectsFunc func() (object.CommitIter, error)

	// CreateBranchFunc mocks the CreateBranch method.
	CreateBranchFunc func(c *config2.Branch) error

	// CreateRemoteFunc mocks the CreateRemote method.
	CreateRemoteFunc func(c *config2.RemoteConfig) (*git.Remote, error)

	// CreateRemoteAnonymousFunc mocks the CreateRemoteAnonymous method.
	CreateRemoteAnonymousFunc func(c *config2.RemoteConfig) (*git.Remote, error)

	// DeleteBranchFunc mocks the DeleteBranch method.
	DeleteBranchFunc func(name string) error

	// DeleteRemoteFunc mocks the DeleteRemote method.
	DeleteRemoteFunc func(name string) error

	// HeadFunc mocks the Head method.
	HeadFunc func() (*plumbing.Reference, error)

	// ObjectFunc mocks the Object method.
	ObjectFunc func(t plumbing.ObjectType, h plumbing.Hash) (object.Object, error)

	// ObjectsFunc mocks the Objects method.
	ObjectsFunc func() (*object.ObjectIter, error)

	// PushFunc mocks the Push method.
	PushFunc func(o *git.PushOptions) error

	// ReferenceFunc mocks the Reference method.
	ReferenceFunc func(name plumbing.ReferenceName, resolved bool) (*plumbing.Reference, error)

	// ReferencesFunc mocks the References method.
	ReferencesFunc func() (storer.ReferenceIter, error)

	// RemoteFunc mocks the Remote method.
	RemoteFunc func(name string) (*git.Remote, error)

	// RemotesFunc mocks the Remotes method.
	RemotesFunc func() ([]*git.Remote, error)

	// ResolveRevisionFunc mocks the ResolveRevision method.
	ResolveRevisionFunc func(rev plumbing.Revision) (*plumbing.Hash, error)

	// TreeObjectFunc mocks the TreeObject method.
	TreeObjectFunc func(h plumbing.Hash) (*object.Tree, error)

	// TreeObjectsFunc mocks the TreeObjects method.
	TreeObjectsFunc func() (*object.TreeIter, error)

	// WorktreeFunc mocks the Worktree method.
	WorktreeFunc func() (*git.Worktree, error)

	// calculateRemoteHeadReferenceFunc mocks the calculateRemoteHeadReference method.
	calculateRemoteHeadReferenceFunc func(spec []config2.RefSpec, resolvedHead *plumbing.Reference) []*plumbing.Reference

	// cloneFunc mocks the clone method.
	cloneFunc func(ctx context.Context, o *git.CloneOptions) error

	// resolveToCommitHashFunc mocks the resolveToCommitHash method.
	resolveToCommitHashFunc func(h plumbing.Hash) (plumbing.Hash, error)

	// calls tracks calls to the methods.
	calls struct {
		// BlobObject holds details about calls to the BlobObject method.
		BlobObject []struct {
			// H is the h argument value.
			H plumbing.Hash
		}
		// BlobObjects holds details about calls to the BlobObjects method.
		BlobObjects []struct {
		}
		// Branch holds details about calls to the Branch method.
		Branch []struct {
			// Name is the name argument value.
			Name string
		}
		// Branches holds details about calls to the Branches method.
		Branches []struct {
		}
		// CommitObject holds details about calls to the CommitObject method.
		CommitObject []struct {
			// H is the h argument value.
			H plumbing.Hash
		}
		// CommitObjects holds details about calls to the CommitObjects method.
		CommitObjects []struct {
		}
		// CreateBranch holds details about calls to the CreateBranch method.
		CreateBranch []struct {
			// C is the c argument value.
			C *config2.Branch
		}
		// CreateRemote holds details about calls to the CreateRemote method.
		CreateRemote []struct {
			// C is the c argument value.
			C *config2.RemoteConfig
		}
		// CreateRemoteAnonymous holds details about calls to the CreateRemoteAnonymous method.
		CreateRemoteAnonymous []struct {
			// C is the c argument value.
			C *config2.RemoteConfig
		}
		// DeleteBranch holds details about calls to the DeleteBranch method.
		DeleteBranch []struct {
			// Name is the name argument value.
			Name string
		}
		// DeleteRemote holds details about calls to the DeleteRemote method.
		DeleteRemote []struct {
			// Name is the name argument value.
			Name string
		}
		// Head holds details about calls to the Head method.
		Head []struct {
		}
		// Object holds details about calls to the Object method.
		Object []struct {
			// T is the t argument value.
			T plumbing.ObjectType
			// H is the h argument value.
			H plumbing.Hash
		}
		// Objects holds details about calls to the Objects method.
		Objects []struct {
		}
		// Push holds details about calls to the Push method.
		Push []struct {
			// O is the o argument value.
			O *git.PushOptions
		}
		// Reference holds details about calls to the Reference method.
		Reference []struct {
			// Name is the name argument value.
			Name plumbing.ReferenceName
			// Resolved is the resolved argument value.
			Resolved bool
		}
		// References holds details about calls to the References method.
		References []struct {
		}
		// Remote holds details about calls to the Remote method.
		Remote []struct {
			// Name is the name argument value.
			Name string
		}
		// Remotes holds details about calls to the Remotes method.
		Remotes []struct {
		}
		// ResolveRevision holds details about calls to the ResolveRevision method.
		ResolveRevision []struct {
			// Rev is the rev argument value.
			Rev plumbing.Revision
		}
		// TreeObject holds details about calls to the TreeObject method.
		TreeObject []struct {
			// H is the h argument value.
			H plumbing.Hash
		}
		// TreeObjects holds details about calls to the TreeObjects method.
		TreeObjects []struct {
		}
		// Worktree holds details about calls to the Worktree method.
		Worktree []struct {
		}
		// calculateRemoteHeadReference holds details about calls to the calculateRemoteHeadReference method.
		calculateRemoteHeadReference []struct {
			// Spec is the spec argument value.
			Spec []config2.RefSpec
			// ResolvedHead is the resolvedHead argument value.
			ResolvedHead *plumbing.Reference
		}
		// clone holds details about calls to the clone method.
		clone []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// O is the o argument value.
			O *git.CloneOptions
		}
		// resolveToCommitHash holds details about calls to the resolveToCommitHash method.
		resolveToCommitHash []struct {
			// H is the h argument value.
			H plumbing.Hash
		}
	}
	lockBlobObject                   sync.RWMutex
	lockBlobObjects                  sync.RWMutex
	lockBranch                       sync.RWMutex
	lockBranches                     sync.RWMutex
	lockCommitObject                 sync.RWMutex
	lockCommitObjects                sync.RWMutex
	lockCreateBranch                 sync.RWMutex
	lockCreateRemote                 sync.RWMutex
	lockCreateRemoteAnonymous        sync.RWMutex
	lockDeleteBranch                 sync.RWMutex
	lockDeleteRemote                 sync.RWMutex
	lockHead                         sync.RWMutex
	lockObject                       sync.RWMutex
	lockObjects                      sync.RWMutex
	lockPush                         sync.RWMutex
	lockReference                    sync.RWMutex
	lockReferences                   sync.RWMutex
	lockRemote                       sync.RWMutex
	lockRemotes                      sync.RWMutex
	lockResolveRevision              sync.RWMutex
	lockTreeObject                   sync.RWMutex
	lockTreeObjects                  sync.RWMutex
	lockWorktree                     sync.RWMutex
	lockcalculateRemoteHeadReference sync.RWMutex
	lockclone                        sync.RWMutex
	lockresolveToCommitHash          sync.RWMutex
}

// BlobObject calls BlobObjectFunc.
func (mock *RepositoryMock) BlobObject(h plumbing.Hash) (*object.Blob, error) {
	if mock.BlobObjectFunc == nil {
		panic("RepositoryMock.BlobObjectFunc: method is nil but Repository.BlobObject was just called")
	}
	callInfo := struct {
		H plumbing.Hash
	}{
		H: h,
	}
	mock.lockBlobObject.Lock()
	mock.calls.BlobObject = append(mock.calls.BlobObject, callInfo)
	mock.lockBlobObject.Unlock()
	return mock.BlobObjectFunc(h)
}

// BlobObjectCalls gets all the calls that were made to BlobObject.
// Check the length with:
//     len(mockedRepository.BlobObjectCalls())
func (mock *RepositoryMock) BlobObjectCalls() []struct {
	H plumbing.Hash
} {
	var calls []struct {
		H plumbing.Hash
	}
	mock.lockBlobObject.RLock()
	calls = mock.calls.BlobObject
	mock.lockBlobObject.RUnlock()
	return calls
}

// BlobObjects calls BlobObjectsFunc.
func (mock *RepositoryMock) BlobObjects() (*object.BlobIter, error) {
	if mock.BlobObjectsFunc == nil {
		panic("RepositoryMock.BlobObjectsFunc: method is nil but Repository.BlobObjects was just called")
	}
	callInfo := struct {
	}{}
	mock.lockBlobObjects.Lock()
	mock.calls.BlobObjects = append(mock.calls.BlobObjects, callInfo)
	mock.lockBlobObjects.Unlock()
	return mock.BlobObjectsFunc()
}

// BlobObjectsCalls gets all the calls that were made to BlobObjects.
// Check the length with:
//     len(mockedRepository.BlobObjectsCalls())
func (mock *RepositoryMock) BlobObjectsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockBlobObjects.RLock()
	calls = mock.calls.BlobObjects
	mock.lockBlobObjects.RUnlock()
	return calls
}

// Branch calls BranchFunc.
func (mock *RepositoryMock) Branch(name string) (*config2.Branch, error) {
	if mock.BranchFunc == nil {
		panic("RepositoryMock.BranchFunc: method is nil but Repository.Branch was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockBranch.Lock()
	mock.calls.Branch = append(mock.calls.Branch, callInfo)
	mock.lockBranch.Unlock()
	return mock.BranchFunc(name)
}

// BranchCalls gets all the calls that were made to Branch.
// Check the length with:
//     len(mockedRepository.BranchCalls())
func (mock *RepositoryMock) BranchCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockBranch.RLock()
	calls = mock.calls.Branch
	mock.lockBranch.RUnlock()
	return calls
}

// Branches calls BranchesFunc.
func (mock *RepositoryMock) Branches() (storer.ReferenceIter, error) {
	if mock.BranchesFunc == nil {
		panic("RepositoryMock.BranchesFunc: method is nil but Repository.Branches was just called")
	}
	callInfo := struct {
	}{}
	mock.lockBranches.Lock()
	mock.calls.Branches = append(mock.calls.Branches, callInfo)
	mock.lockBranches.Unlock()
	return mock.BranchesFunc()
}

// BranchesCalls gets all the calls that were made to Branches.
// Check the length with:
//     len(mockedRepository.BranchesCalls())
func (mock *RepositoryMock) BranchesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockBranches.RLock()
	calls = mock.calls.Branches
	mock.lockBranches.RUnlock()
	return calls
}

// CommitObject calls CommitObjectFunc.
func (mock *RepositoryMock) CommitObject(h plumbing.Hash) (*object.Commit, error) {
	if mock.CommitObjectFunc == nil {
		panic("RepositoryMock.CommitObjectFunc: method is nil but Repository.CommitObject was just called")
	}
	callInfo := struct {
		H plumbing.Hash
	}{
		H: h,
	}
	mock.lockCommitObject.Lock()
	mock.calls.CommitObject = append(mock.calls.CommitObject, callInfo)
	mock.lockCommitObject.Unlock()
	return mock.CommitObjectFunc(h)
}

// CommitObjectCalls gets all the calls that were made to CommitObject.
// Check the length with:
//     len(mockedRepository.CommitObjectCalls())
func (mock *RepositoryMock) CommitObjectCalls() []struct {
	H plumbing.Hash
} {
	var calls []struct {
		H plumbing.Hash
	}
	mock.lockCommitObject.RLock()
	calls = mock.calls.CommitObject
	mock.lockCommitObject.RUnlock()
	return calls
}

// CommitObjects calls CommitObjectsFunc.
func (mock *RepositoryMock) CommitObjects() (object.CommitIter, error) {
	if mock.CommitObjectsFunc == nil {
		panic("RepositoryMock.CommitObjectsFunc: method is nil but Repository.CommitObjects was just called")
	}
	callInfo := struct {
	}{}
	mock.lockCommitObjects.Lock()
	mock.calls.CommitObjects = append(mock.calls.CommitObjects, callInfo)
	mock.lockCommitObjects.Unlock()
	return mock.CommitObjectsFunc()
}

// CommitObjectsCalls gets all the calls that were made to CommitObjects.
// Check the length with:
//     len(mockedRepository.CommitObjectsCalls())
func (mock *RepositoryMock) CommitObjectsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockCommitObjects.RLock()
	calls = mock.calls.CommitObjects
	mock.lockCommitObjects.RUnlock()
	return calls
}

// CreateBranch calls CreateBranchFunc.
func (mock *RepositoryMock) CreateBranch(c *config2.Branch) error {
	if mock.CreateBranchFunc == nil {
		panic("RepositoryMock.CreateBranchFunc: method is nil but Repository.CreateBranch was just called")
	}
	callInfo := struct {
		C *config2.Branch
	}{
		C: c,
	}
	mock.lockCreateBranch.Lock()
	mock.calls.CreateBranch = append(mock.calls.CreateBranch, callInfo)
	mock.lockCreateBranch.Unlock()
	return mock.CreateBranchFunc(c)
}

// CreateBranchCalls gets all the calls that were made to CreateBranch.
// Check the length with:
//     len(mockedRepository.CreateBranchCalls())
func (mock *RepositoryMock) CreateBranchCalls() []struct {
	C *config2.Branch
} {
	var calls []struct {
		C *config2.Branch
	}
	mock.lockCreateBranch.RLock()
	calls = mock.calls.CreateBranch
	mock.lockCreateBranch.RUnlock()
	return calls
}

// CreateRemote calls CreateRemoteFunc.
func (mock *RepositoryMock) CreateRemote(c *config2.RemoteConfig) (*git.Remote, error) {
	if mock.CreateRemoteFunc == nil {
		panic("RepositoryMock.CreateRemoteFunc: method is nil but Repository.CreateRemote was just called")
	}
	callInfo := struct {
		C *config2.RemoteConfig
	}{
		C: c,
	}
	mock.lockCreateRemote.Lock()
	mock.calls.CreateRemote = append(mock.calls.CreateRemote, callInfo)
	mock.lockCreateRemote.Unlock()
	return mock.CreateRemoteFunc(c)
}

// CreateRemoteCalls gets all the calls that were made to CreateRemote.
// Check the length with:
//     len(mockedRepository.CreateRemoteCalls())
func (mock *RepositoryMock) CreateRemoteCalls() []struct {
	C *config2.RemoteConfig
} {
	var calls []struct {
		C *config2.RemoteConfig
	}
	mock.lockCreateRemote.RLock()
	calls = mock.calls.CreateRemote
	mock.lockCreateRemote.RUnlock()
	return calls
}

// CreateRemoteAnonymous calls CreateRemoteAnonymousFunc.
func (mock *RepositoryMock) CreateRemoteAnonymous(c *config2.RemoteConfig) (*git.Remote, error) {
	if mock.CreateRemoteAnonymousFunc == nil {
		panic("RepositoryMock.CreateRemoteAnonymousFunc: method is nil but Repository.CreateRemoteAnonymous was just called")
	}
	callInfo := struct {
		C *config2.RemoteConfig
	}{
		C: c,
	}
	mock.lockCreateRemoteAnonymous.Lock()
	mock.calls.CreateRemoteAnonymous = append(mock.calls.CreateRemoteAnonymous, callInfo)
	mock.lockCreateRemoteAnonymous.Unlock()
	return mock.CreateRemoteAnonymousFunc(c)
}

// CreateRemoteAnonymousCalls gets all the calls that were made to CreateRemoteAnonymous.
// Check the length with:
//     len(mockedRepository.CreateRemoteAnonymousCalls())
func (mock *RepositoryMock) CreateRemoteAnonymousCalls() []struct {
	C *config2.RemoteConfig
} {
	var calls []struct {
		C *config2.RemoteConfig
	}
	mock.lockCreateRemoteAnonymous.RLock()
	calls = mock.calls.CreateRemoteAnonymous
	mock.lockCreateRemoteAnonymous.RUnlock()
	return calls
}

// DeleteBranch calls DeleteBranchFunc.
func (mock *RepositoryMock) DeleteBranch(name string) error {
	if mock.DeleteBranchFunc == nil {
		panic("RepositoryMock.DeleteBranchFunc: method is nil but Repository.DeleteBranch was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockDeleteBranch.Lock()
	mock.calls.DeleteBranch = append(mock.calls.DeleteBranch, callInfo)
	mock.lockDeleteBranch.Unlock()
	return mock.DeleteBranchFunc(name)
}

// DeleteBranchCalls gets all the calls that were made to DeleteBranch.
// Check the length with:
//     len(mockedRepository.DeleteBranchCalls())
func (mock *RepositoryMock) DeleteBranchCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockDeleteBranch.RLock()
	calls = mock.calls.DeleteBranch
	mock.lockDeleteBranch.RUnlock()
	return calls
}

// DeleteRemote calls DeleteRemoteFunc.
func (mock *RepositoryMock) DeleteRemote(name string) error {
	if mock.DeleteRemoteFunc == nil {
		panic("RepositoryMock.DeleteRemoteFunc: method is nil but Repository.DeleteRemote was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockDeleteRemote.Lock()
	mock.calls.DeleteRemote = append(mock.calls.DeleteRemote, callInfo)
	mock.lockDeleteRemote.Unlock()
	return mock.DeleteRemoteFunc(name)
}

// DeleteRemoteCalls gets all the calls that were made to DeleteRemote.
// Check the length with:
//     len(mockedRepository.DeleteRemoteCalls())
func (mock *RepositoryMock) DeleteRemoteCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockDeleteRemote.RLock()
	calls = mock.calls.DeleteRemote
	mock.lockDeleteRemote.RUnlock()
	return calls
}

// Head calls HeadFunc.
func (mock *RepositoryMock) Head() (*plumbing.Reference, error) {
	if mock.HeadFunc == nil {
		panic("RepositoryMock.HeadFunc: method is nil but Repository.Head was just called")
	}
	callInfo := struct {
	}{}
	mock.lockHead.Lock()
	mock.calls.Head = append(mock.calls.Head, callInfo)
	mock.lockHead.Unlock()
	return mock.HeadFunc()
}

// HeadCalls gets all the calls that were made to Head.
// Check the length with:
//     len(mockedRepository.HeadCalls())
func (mock *RepositoryMock) HeadCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockHead.RLock()
	calls = mock.calls.Head
	mock.lockHead.RUnlock()
	return calls
}

// Object calls ObjectFunc.
func (mock *RepositoryMock) Object(t plumbing.ObjectType, h plumbing.Hash) (object.Object, error) {
	if mock.ObjectFunc == nil {
		panic("RepositoryMock.ObjectFunc: method is nil but Repository.Object was just called")
	}
	callInfo := struct {
		T plumbing.ObjectType
		H plumbing.Hash
	}{
		T: t,
		H: h,
	}
	mock.lockObject.Lock()
	mock.calls.Object = append(mock.calls.Object, callInfo)
	mock.lockObject.Unlock()
	return mock.ObjectFunc(t, h)
}

// ObjectCalls gets all the calls that were made to Object.
// Check the length with:
//     len(mockedRepository.ObjectCalls())
func (mock *RepositoryMock) ObjectCalls() []struct {
	T plumbing.ObjectType
	H plumbing.Hash
} {
	var calls []struct {
		T plumbing.ObjectType
		H plumbing.Hash
	}
	mock.lockObject.RLock()
	calls = mock.calls.Object
	mock.lockObject.RUnlock()
	return calls
}

// Objects calls ObjectsFunc.
func (mock *RepositoryMock) Objects() (*object.ObjectIter, error) {
	if mock.ObjectsFunc == nil {
		panic("RepositoryMock.ObjectsFunc: method is nil but Repository.Objects was just called")
	}
	callInfo := struct {
	}{}
	mock.lockObjects.Lock()
	mock.calls.Objects = append(mock.calls.Objects, callInfo)
	mock.lockObjects.Unlock()
	return mock.ObjectsFunc()
}

// ObjectsCalls gets all the calls that were made to Objects.
// Check the length with:
//     len(mockedRepository.ObjectsCalls())
func (mock *RepositoryMock) ObjectsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockObjects.RLock()
	calls = mock.calls.Objects
	mock.lockObjects.RUnlock()
	return calls
}

// Push calls PushFunc.
func (mock *RepositoryMock) Push(o *git.PushOptions) error {
	if mock.PushFunc == nil {
		panic("RepositoryMock.PushFunc: method is nil but Repository.Push was just called")
	}
	callInfo := struct {
		O *git.PushOptions
	}{
		O: o,
	}
	mock.lockPush.Lock()
	mock.calls.Push = append(mock.calls.Push, callInfo)
	mock.lockPush.Unlock()
	return mock.PushFunc(o)
}

// PushCalls gets all the calls that were made to Push.
// Check the length with:
//     len(mockedRepository.PushCalls())
func (mock *RepositoryMock) PushCalls() []struct {
	O *git.PushOptions
} {
	var calls []struct {
		O *git.PushOptions
	}
	mock.lockPush.RLock()
	calls = mock.calls.Push
	mock.lockPush.RUnlock()
	return calls
}

// Reference calls ReferenceFunc.
func (mock *RepositoryMock) Reference(name plumbing.ReferenceName, resolved bool) (*plumbing.Reference, error) {
	if mock.ReferenceFunc == nil {
		panic("RepositoryMock.ReferenceFunc: method is nil but Repository.Reference was just called")
	}
	callInfo := struct {
		Name     plumbing.ReferenceName
		Resolved bool
	}{
		Name:     name,
		Resolved: resolved,
	}
	mock.lockReference.Lock()
	mock.calls.Reference = append(mock.calls.Reference, callInfo)
	mock.lockReference.Unlock()
	return mock.ReferenceFunc(name, resolved)
}

// ReferenceCalls gets all the calls that were made to Reference.
// Check the length with:
//     len(mockedRepository.ReferenceCalls())
func (mock *RepositoryMock) ReferenceCalls() []struct {
	Name     plumbing.ReferenceName
	Resolved bool
} {
	var calls []struct {
		Name     plumbing.ReferenceName
		Resolved bool
	}
	mock.lockReference.RLock()
	calls = mock.calls.Reference
	mock.lockReference.RUnlock()
	return calls
}

// References calls ReferencesFunc.
func (mock *RepositoryMock) References() (storer.ReferenceIter, error) {
	if mock.ReferencesFunc == nil {
		panic("RepositoryMock.ReferencesFunc: method is nil but Repository.References was just called")
	}
	callInfo := struct {
	}{}
	mock.lockReferences.Lock()
	mock.calls.References = append(mock.calls.References, callInfo)
	mock.lockReferences.Unlock()
	return mock.ReferencesFunc()
}

// ReferencesCalls gets all the calls that were made to References.
// Check the length with:
//     len(mockedRepository.ReferencesCalls())
func (mock *RepositoryMock) ReferencesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockReferences.RLock()
	calls = mock.calls.References
	mock.lockReferences.RUnlock()
	return calls
}

// Remote calls RemoteFunc.
func (mock *RepositoryMock) Remote(name string) (*git.Remote, error) {
	if mock.RemoteFunc == nil {
		panic("RepositoryMock.RemoteFunc: method is nil but Repository.Remote was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockRemote.Lock()
	mock.calls.Remote = append(mock.calls.Remote, callInfo)
	mock.lockRemote.Unlock()
	return mock.RemoteFunc(name)
}

// RemoteCalls gets all the calls that were made to Remote.
// Check the length with:
//     len(mockedRepository.RemoteCalls())
func (mock *RepositoryMock) RemoteCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockRemote.RLock()
	calls = mock.calls.Remote
	mock.lockRemote.RUnlock()
	return calls
}

// Remotes calls RemotesFunc.
func (mock *RepositoryMock) Remotes() ([]*git.Remote, error) {
	if mock.RemotesFunc == nil {
		panic("RepositoryMock.RemotesFunc: method is nil but Repository.Remotes was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRemotes.Lock()
	mock.calls.Remotes = append(mock.calls.Remotes, callInfo)
	mock.lockRemotes.Unlock()
	return mock.RemotesFunc()
}

// RemotesCalls gets all the calls that were made to Remotes.
// Check the length with:
//     len(mockedRepository.RemotesCalls())
func (mock *RepositoryMock) RemotesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRemotes.RLock()
	calls = mock.calls.Remotes
	mock.lockRemotes.RUnlock()
	return calls
}

// ResolveRevision calls ResolveRevisionFunc.
func (mock *RepositoryMock) ResolveRevision(rev plumbing.Revision) (*plumbing.Hash, error) {
	if mock.ResolveRevisionFunc == nil {
		panic("RepositoryMock.ResolveRevisionFunc: method is nil but Repository.ResolveRevision was just called")
	}
	callInfo := struct {
		Rev plumbing.Revision
	}{
		Rev: rev,
	}
	mock.lockResolveRevision.Lock()
	mock.calls.ResolveRevision = append(mock.calls.ResolveRevision, callInfo)
	mock.lockResolveRevision.Unlock()
	return mock.ResolveRevisionFunc(rev)
}

// ResolveRevisionCalls gets all the calls that were made to ResolveRevision.
// Check the length with:
//     len(mockedRepository.ResolveRevisionCalls())
func (mock *RepositoryMock) ResolveRevisionCalls() []struct {
	Rev plumbing.Revision
} {
	var calls []struct {
		Rev plumbing.Revision
	}
	mock.lockResolveRevision.RLock()
	calls = mock.calls.ResolveRevision
	mock.lockResolveRevision.RUnlock()
	return calls
}

// TreeObject calls TreeObjectFunc.
func (mock *RepositoryMock) TreeObject(h plumbing.Hash) (*object.Tree, error) {
	if mock.TreeObjectFunc == nil {
		panic("RepositoryMock.TreeObjectFunc: method is nil but Repository.TreeObject was just called")
	}
	callInfo := struct {
		H plumbing.Hash
	}{
		H: h,
	}
	mock.lockTreeObject.Lock()
	mock.calls.TreeObject = append(mock.calls.TreeObject, callInfo)
	mock.lockTreeObject.Unlock()
	return mock.TreeObjectFunc(h)
}

// TreeObjectCalls gets all the calls that were made to TreeObject.
// Check the length with:
//     len(mockedRepository.TreeObjectCalls())
func (mock *RepositoryMock) TreeObjectCalls() []struct {
	H plumbing.Hash
} {
	var calls []struct {
		H plumbing.Hash
	}
	mock.lockTreeObject.RLock()
	calls = mock.calls.TreeObject
	mock.lockTreeObject.RUnlock()
	return calls
}

// TreeObjects calls TreeObjectsFunc.
func (mock *RepositoryMock) TreeObjects() (*object.TreeIter, error) {
	if mock.TreeObjectsFunc == nil {
		panic("RepositoryMock.TreeObjectsFunc: method is nil but Repository.TreeObjects was just called")
	}
	callInfo := struct {
	}{}
	mock.lockTreeObjects.Lock()
	mock.calls.TreeObjects = append(mock.calls.TreeObjects, callInfo)
	mock.lockTreeObjects.Unlock()
	return mock.TreeObjectsFunc()
}

// TreeObjectsCalls gets all the calls that were made to TreeObjects.
// Check the length with:
//     len(mockedRepository.TreeObjectsCalls())
func (mock *RepositoryMock) TreeObjectsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTreeObjects.RLock()
	calls = mock.calls.TreeObjects
	mock.lockTreeObjects.RUnlock()
	return calls
}

// Worktree calls WorktreeFunc.
func (mock *RepositoryMock) Worktree() (*git.Worktree, error) {
	if mock.WorktreeFunc == nil {
		panic("RepositoryMock.WorktreeFunc: method is nil but Repository.Worktree was just called")
	}
	callInfo := struct {
	}{}
	mock.lockWorktree.Lock()
	mock.calls.Worktree = append(mock.calls.Worktree, callInfo)
	mock.lockWorktree.Unlock()
	return mock.WorktreeFunc()
}

// WorktreeCalls gets all the calls that were made to Worktree.
// Check the length with:
//     len(mockedRepository.WorktreeCalls())
func (mock *RepositoryMock) WorktreeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockWorktree.RLock()
	calls = mock.calls.Worktree
	mock.lockWorktree.RUnlock()
	return calls
}

// calculateRemoteHeadReference calls calculateRemoteHeadReferenceFunc.
func (mock *RepositoryMock) calculateRemoteHeadReference(spec []config2.RefSpec, resolvedHead *plumbing.Reference) []*plumbing.Reference {
	if mock.calculateRemoteHeadReferenceFunc == nil {
		panic("RepositoryMock.calculateRemoteHeadReferenceFunc: method is nil but Repository.calculateRemoteHeadReference was just called")
	}
	callInfo := struct {
		Spec         []config2.RefSpec
		ResolvedHead *plumbing.Reference
	}{
		Spec:         spec,
		ResolvedHead: resolvedHead,
	}
	mock.lockcalculateRemoteHeadReference.Lock()
	mock.calls.calculateRemoteHeadReference = append(mock.calls.calculateRemoteHeadReference, callInfo)
	mock.lockcalculateRemoteHeadReference.Unlock()
	return mock.calculateRemoteHeadReferenceFunc(spec, resolvedHead)
}

// calculateRemoteHeadReferenceCalls gets all the calls that were made to calculateRemoteHeadReference.
// Check the length with:
//     len(mockedRepository.calculateRemoteHeadReferenceCalls())
func (mock *RepositoryMock) calculateRemoteHeadReferenceCalls() []struct {
	Spec         []config2.RefSpec
	ResolvedHead *plumbing.Reference
} {
	var calls []struct {
		Spec         []config2.RefSpec
		ResolvedHead *plumbing.Reference
	}
	mock.lockcalculateRemoteHeadReference.RLock()
	calls = mock.calls.calculateRemoteHeadReference
	mock.lockcalculateRemoteHeadReference.RUnlock()
	return calls
}

// clone calls cloneFunc.
func (mock *RepositoryMock) clone(ctx context.Context, o *git.CloneOptions) error {
	if mock.cloneFunc == nil {
		panic("RepositoryMock.cloneFunc: method is nil but Repository.clone was just called")
	}
	callInfo := struct {
		Ctx context.Context
		O   *git.CloneOptions
	}{
		Ctx: ctx,
		O:   o,
	}
	mock.lockclone.Lock()
	mock.calls.clone = append(mock.calls.clone, callInfo)
	mock.lockclone.Unlock()
	return mock.cloneFunc(ctx, o)
}

// cloneCalls gets all the calls that were made to clone.
// Check the length with:
//     len(mockedRepository.cloneCalls())
func (mock *RepositoryMock) cloneCalls() []struct {
	Ctx context.Context
	O   *git.CloneOptions
} {
	var calls []struct {
		Ctx context.Context
		O   *git.CloneOptions
	}
	mock.lockclone.RLock()
	calls = mock.calls.clone
	mock.lockclone.RUnlock()
	return calls
}

// resolveToCommitHash calls resolveToCommitHashFunc.
func (mock *RepositoryMock) resolveToCommitHash(h plumbing.Hash) (plumbing.Hash, error) {
	if mock.resolveToCommitHashFunc == nil {
		panic("RepositoryMock.resolveToCommitHashFunc: method is nil but Repository.resolveToCommitHash was just called")
	}
	callInfo := struct {
		H plumbing.Hash
	}{
		H: h,
	}
	mock.lockresolveToCommitHash.Lock()
	mock.calls.resolveToCommitHash = append(mock.calls.resolveToCommitHash, callInfo)
	mock.lockresolveToCommitHash.Unlock()
	return mock.resolveToCommitHashFunc(h)
}

// resolveToCommitHashCalls gets all the calls that were made to resolveToCommitHash.
// Check the length with:
//     len(mockedRepository.resolveToCommitHashCalls())
func (mock *RepositoryMock) resolveToCommitHashCalls() []struct {
	H plumbing.Hash
} {
	var calls []struct {
		H plumbing.Hash
	}
	mock.lockresolveToCommitHash.RLock()
	calls = mock.calls.resolveToCommitHash
	mock.lockresolveToCommitHash.RUnlock()
	return calls
}