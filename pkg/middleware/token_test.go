/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"iam/pkg/util"
)

func TestTokenAuth(t *testing.T) {
	t.Parallel()

	// 1. right
	w := httptest.NewRecorder()
	c := util.CreateTestContextWithDefaultRequest(w)

	q := c.Request.URL.Query()
	q.Add("token", "test")
	c.Request.URL.RawQuery = q.Encode()

	TokenAuth("test")(c)

	assert.Equal(t, 200, w.Code)

	// 2. right
	w = httptest.NewRecorder()
	c = util.CreateTestContextWithDefaultRequest(w)

	TokenAuth("test")(c)

	assert.Equal(t, 401, w.Code)
}
