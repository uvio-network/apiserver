package userstorage

// TODO cut off the user IDs with the lowest competence scores
func (r *Redigo) DeleteCompetence() error {
	return nil
}

// TODO should DeleteCompetence and DeleteIntegrity be one method? how do we search for users?

// TODO cut off the user IDs with the lowest integrity scores
func (r *Redigo) DeleteIntegrity() error {
	return nil
}
