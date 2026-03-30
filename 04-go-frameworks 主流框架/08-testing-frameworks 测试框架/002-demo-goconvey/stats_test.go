package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStats(t *testing.T) {
	Convey("Given a slice of scores", t, func() {
		scores := []int{80, 90, 100}

		Convey("Average should be correct", func() {
			So(Average(scores), ShouldEqual, 90)
		})

		Convey("Max should be correct", func() {
			So(Max(scores), ShouldEqual, 100)
		})
	})
}
