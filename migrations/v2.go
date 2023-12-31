/*
 * Copyright © 2019 Musing Studio LLC.
 *
 * This file is part of WriteFreely.
 *
 * WriteFreely is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package migrations

func supportInstancePages(db *datastore) error {
	t, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = t.Exec(`ALTER TABLE appcontent ADD COLUMN title ` + db.typeVarChar(255) + db.collateMultiByte() + ` NULL`)
	if err != nil {
		t.Rollback()
		return err
	}

	_, err = t.Exec(`ALTER TABLE appcontent ADD COLUMN content_type ` + db.typeVarChar(36) + ` DEFAULT 'page' NOT NULL`)
	if err != nil {
		t.Rollback()
		return err
	}

	err = t.Commit()
	if err != nil {
		t.Rollback()
		return err
	}

	return nil
}
