import "time"

func (o *RPC) SignIn(args *data_define.SignInExArgs, reply *data_define.SignInExReply) (err error) {
	redisConn := db.NewDBProxy().RedisPool().Get()
	defer redisConn.Close()

	redisConn.Do("SELECT", data_define.RedisOther)

	if reply.Status, err = redis.Uint64(redisConn.Do("HGET", "signin_status", args.UserID)); err != nil && err != redis.ErrNil {
		return err
	}

	now := time.Now()
	day := now.Day()
	mask := uint64(1) << uint(day-1)

	// 已经签到
	if reply.Status&mask != 0 {
		// 返回客户端错误码
		return data_define.ErrAlreadySignIn
	}

	reply.Status |= mask

	// 存储到Redis中
	if _, err = redisConn.Do("HSET", "signin_status", args.UserID, reply.Status); err != nil {
		return err
	}

	// 抽奖次数增加一次
	if _, err = redisConn.Do("HINCRBY", "draw_count", args.UserID, 1); err != nil {
		return err
	}
	return nil
}