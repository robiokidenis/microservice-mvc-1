/*
 * Copyright (c) 2021
 * Created at 5/26/21 1:49 AM
 * Created by robiokidenis
 * Email robiokidenis@gmail.com
 * Github https://github.com/robiokidenis
 */

package auth

type authRequest struct {
	SocialType  string `json:"social_type" binding:"required"`
	SocialToken string `json:"social_token" binding:"required"`
}

func (req authRequest) GetSocialType() string {
	return req.SocialType
}

func (req authRequest) GetSocialToken() string {
	return req.SocialToken
}
