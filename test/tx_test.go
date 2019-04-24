/*
 * Copyright 2019 ObjectBox Ltd. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package objectbox_test

import (
	"errors"
	"testing"

	"github.com/objectbox/objectbox-go/test/assert"
	"github.com/objectbox/objectbox-go/test/model/iot"
)

func TestTransactionMassiveInsert(t *testing.T) {
	ob := iot.LoadEmptyTestObjectBox()
	defer ob.Close()

	var box = iot.BoxForEvent(ob)

	assert.NoErr(t, box.RemoveAll())

	// TODO increase this after we switch to the c-api box interface,
	//  current implementation still creates transaction for each Put()
	var insert = uint64(10000)

	err := ob.Update(func() error {
		for i := insert; i > 0; i-- {
			_, err := box.Put(&iot.Event{})
			assert.NoErr(t, err)
		}
		return nil
	})

	count, err := box.Count()
	assert.NoErr(t, err)
	assert.Eq(t, insert, count)
}

func TestTransactionRollback(t *testing.T) {
	ob := iot.LoadEmptyTestObjectBox()
	defer ob.Close()

	var box = iot.BoxForEvent(ob)

	assert.NoErr(t, box.RemoveAll())

	var insert = make([]*iot.Event, 100)
	for i := 0; i < len(insert); i++ {
		insert[i] = &iot.Event{}
	}

	_, err := box.PutAll(insert)
	assert.NoErr(t, err)

	count, err := box.Count()
	assert.NoErr(t, err)
	assert.Eq(t, len(insert), int(count))

	// rolled-back Tx
	var expected = errors.New("expected")
	assert.Eq(t, expected, ob.Update(func() error {
		assert.NoErr(t, box.RemoveAll())
		return expected
	}))

	count, err = box.Count()
	assert.NoErr(t, err)
	assert.Eq(t, len(insert), int(count))

	// successful tx
	assert.NoErr(t, ob.Update(func() error {
		assert.NoErr(t, box.RemoveAll())
		return nil
	}))

	count, err = box.Count()
	assert.NoErr(t, err)
	assert.Eq(t, 0, int(count))

}
