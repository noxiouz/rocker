/*-
 * Copyright 2015 Grammarly, Inc.
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

package build2

import "io"

// readerVoidCloser is a hack of the improved go-dockerclient's hijacking behavior
// It simply wraps io.Reader (os.Stdin in our case) and discards any Close() call.
//
// It's important because we don't want to close os.Stdin for two reasons:
// 1. We need to restore the terminal back from the raw mode after ATTACH
// 2. There can be other ATTACH instructions for which we need an open stdin
//
// See additional notes in the runContainerAttachStdin() function
type readerVoidCloser struct {
	reader io.Reader
}

// Read reads from current reader
func (r readerVoidCloser) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

// Close is a viod function, does nothing
func (r readerVoidCloser) Close() error {
	return nil
}
