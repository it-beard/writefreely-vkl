/*
 * Copyright © 2020 Musing Studio LLC.
 *
 * This file is part of WriteFreely.
 *
 * WriteFreely is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package migrations

func changeCollectionDescriptionFieldLength(db *datastore) error {
	t, err := db.Begin()
	if err != nil {
		t.Rollback()
		return err
	}

	if db.driverName == driverMySQL {
		_, err = t.Exec(`ALTER TABLE collections MODIFY description VARCHAR(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;`)
		if err != nil {
			t.Rollback()
			return err
		}
	} else {
		// Stab for SQLite
		_, err = t.Exec(`SELECT 1;`)
		if err != nil {
			t.Rollback()
			return err
		}
	}

	err = t.Commit()
	if err != nil {
		t.Rollback()
		return err
	}

	return nil
}
