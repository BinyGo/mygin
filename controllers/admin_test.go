/*
 * @Auther: Biny
 * @Description:
 * @Date: 2022-01-22 20:37:41
 * @LastEditTime: 2022-01-24 22:23:00
 */
package controllers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAdminAuth(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AdminAuth(tt.args.c)
		})
	}
}
