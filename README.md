```
Client 				 								                                                    Auth server
1. user login								                                ---->			         generate {token, refresh_token}
2. receive {token, refresh_token}					                        <----
3. use token to request other server (posts, comments,.... servers)
4. when token expired, use refresh_token to request to auth server	        ---->			         revoke old-{token, refresh_token}
												                                                     generate new-{token, refresh_token}
5. continue as step 2							                            <----

(when refresh_token is also expried (Ex: user do not interact for too long), requires re-login)
```

References:
Purpose of refresh_token


[Register]
- OTP
- check if email and phone exist
- If email or phone is exist -> do not allow

[Logout]
- revoke token

[Api for admin]

[Forget password]