package tree

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petethepig/pyroscope/pkg/storage/dict"
)

func randStr() []byte {
	buf := make([]byte, 10)
	for i := 0; i < 10; i++ {
		buf[i] = byte(97) + byte(rand.Uint32()%10)
	}
	return buf
}

var _ = Describe("tree package", func() {
	Context("tree", func() {
		d := dict.New()

		tree := New()
		tree.Insert([]byte("a;b"), uint64(1))
		tree.Insert([]byte("a;c"), uint64(2))

		It("returns correct results", func() {
			Expect(tree.root.childrenNodes).To(HaveLen(1))
			Expect(tree.root.childrenNodes[0].childrenNodes).To(HaveLen(2))
			// Expect(tree.root.childrenNodes[0].label).To(Equal([]byte("a")))
			Expect(tree.root.childrenNodes[0].self).To(Equal(uint64(0)))
			Expect(tree.root.childrenNodes[0].cum).To(Equal(uint64(3)))
			// Expect(tree.root.childrenNodes[0].childrenLabels[0]).To(Equal([]byte("b")))
			// Expect(tree.root.childrenNodes[0].childrenLabels[1]).To(Equal([]byte("c")))
			Expect(tree.root.childrenNodes[0].childrenNodes[0].self).To(Equal(uint64(1)))
			Expect(tree.root.childrenNodes[0].childrenNodes[1].self).To(Equal(uint64(2)))
			Expect(tree.root.childrenNodes[0].childrenNodes[0].cum).To(Equal(uint64(1)))
			Expect(tree.root.childrenNodes[0].childrenNodes[1].cum).To(Equal(uint64(2)))
			Expect(tree.String(d)).To(Equal("\"a;b\" 1\n\"a;c\" 2\n"))
		})
	})

	Context("tree.Merge", func() {
		d := dict.New()

		tree := New()
		tree.Insert([]byte("a;b"), uint64(1))
		tree.Insert([]byte("a;c"), uint64(2))

		tree2 := New()
		tree2.Insert([]byte("a;b"), uint64(1))
		tree2.Insert([]byte("a;c"), uint64(2))

		It("returns correct results", func() {
			Expect(tree.root.childrenNodes).To(HaveLen(1))
			Expect(tree.root.childrenNodes[0].childrenNodes).To(HaveLen(2))
			// Expect(tree.root.childrenNodes[0].label).To(Equal([]byte("a")))
			Expect(tree.root.childrenNodes[0].self).To(Equal(uint64(0)))
			Expect(tree.root.childrenNodes[0].cum).To(Equal(uint64(3)))
			// Expect(tree.root.childrenNodes[0].childrenLabels[0]).To(Equal([]byte("b")))
			// Expect(tree.root.childrenNodes[0].childrenLabels[1]).To(Equal([]byte("c")))
			Expect(tree.root.childrenNodes[0].childrenNodes[0].self).To(Equal(uint64(1)))
			Expect(tree.root.childrenNodes[0].childrenNodes[1].self).To(Equal(uint64(2)))
			Expect(tree.root.childrenNodes[0].childrenNodes[0].cum).To(Equal(uint64(1)))
			Expect(tree.root.childrenNodes[0].childrenNodes[1].cum).To(Equal(uint64(2)))
			Expect(tree.String(d)).To(Equal("\"a;b\" 1\n\"a;c\" 2\n"))
		})
	})
})