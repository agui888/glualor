local cookie_middleware = function(cookieConfig)
    return function(req, res, next)
            req.cookie = {
                set = function(cookie)
                    gluaweb.SetCookie(cookie)
                end,

                get = function (name) 
                    local cookie, err = gluaweb.Cookie(name)
                    if err != nil then
                        print(err)
                        return
                    end
                    return cookie
                end,

                get_all = function ()
                    return gluaweb.Cookies()
                end
            }
        end

        next()
    end
end

return cookie_middleware