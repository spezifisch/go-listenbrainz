// Copyright (C) 2020 Pascal Below (spezifisch)

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package listenbrainz

import (
	"testing"
)

func TestGetListenHistory(t *testing.T) {
	api := API{
		URL:   "http://127.0.0.1:0",
		Token: "baz",
	}

	resp, err := api.GetListenHistory("bob")
	if err == nil {
		t.Error("expected error")
	}
	if resp != nil {
		t.Error("expected nil response")
	}
}
