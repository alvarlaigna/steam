// DON'T EDIT *** generated by scaneo *** DON'T EDIT //

package data

import "database/sql"

func ScanMeta(r *sql.Row) (Meta, error) {
	var s Meta
	if err := r.Scan(
		&s.Id,
		&s.Key,
		&s.Value,
	); err != nil {
		return Meta{}, err
	}
	return s, nil
}

func ScanMetas(rs *sql.Rows) ([]Meta, error) {
	structs := make([]Meta, 0, 16)
	var err error
	for rs.Next() {
		var s Meta
		if err = rs.Scan(
			&s.Id,
			&s.Key,
			&s.Value,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanPermission(r *sql.Row) (Permission, error) {
	var s Permission
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Description,
	); err != nil {
		return Permission{}, err
	}
	return s, nil
}

func ScanPermissions(rs *sql.Rows) ([]Permission, error) {
	structs := make([]Permission, 0, 16)
	var err error
	for rs.Next() {
		var s Permission
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Description,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanEntityType(r *sql.Row) (EntityType, error) {
	var s EntityType
	if err := r.Scan(
		&s.Id,
		&s.Name,
	); err != nil {
		return EntityType{}, err
	}
	return s, nil
}

func ScanEntityTypes(rs *sql.Rows) ([]EntityType, error) {
	structs := make([]EntityType, 0, 16)
	var err error
	for rs.Next() {
		var s EntityType
		if err = rs.Scan(
			&s.Id,
			&s.Name,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanRole(r *sql.Row) (Role, error) {
	var s Role
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Description,
		&s.Created,
	); err != nil {
		return Role{}, err
	}
	return s, nil
}

func ScanRoles(rs *sql.Rows) ([]Role, error) {
	structs := make([]Role, 0, 16)
	var err error
	for rs.Next() {
		var s Role
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Description,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanWorkgroup(r *sql.Row) (Workgroup, error) {
	var s Workgroup
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Description,
		&s.Created,
	); err != nil {
		return Workgroup{}, err
	}
	return s, nil
}

func ScanWorkgroups(rs *sql.Rows) ([]Workgroup, error) {
	structs := make([]Workgroup, 0, 16)
	var err error
	for rs.Next() {
		var s Workgroup
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Description,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanIdentity(r *sql.Row) (Identity, error) {
	var s Identity
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.IsActive,
		&s.LastLogin,
		&s.Created,
	); err != nil {
		return Identity{}, err
	}
	return s, nil
}

func ScanIdentitys(rs *sql.Rows) ([]Identity, error) {
	structs := make([]Identity, 0, 16)
	var err error
	for rs.Next() {
		var s Identity
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.IsActive,
			&s.LastLogin,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanIdentityAndPassword(r *sql.Row) (IdentityAndPassword, error) {
	var s IdentityAndPassword
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Password,
		&s.IsActive,
		&s.LastLogin,
		&s.Created,
	); err != nil {
		return IdentityAndPassword{}, err
	}
	return s, nil
}

func ScanIdentityAndPasswords(rs *sql.Rows) ([]IdentityAndPassword, error) {
	structs := make([]IdentityAndPassword, 0, 16)
	var err error
	for rs.Next() {
		var s IdentityAndPassword
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Password,
			&s.IsActive,
			&s.LastLogin,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

