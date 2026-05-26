package cache

import "github.com/bsm/redislock"

var Locker *redislock.Client

const (
	USER_EDIT_PROFILE_LOCK = "user_profile_lock"
)

func UserEditProfileLockKey(userName string) string {
	return USER_EDIT_PROFILE_LOCK + userName
}
