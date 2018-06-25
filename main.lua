local lor = require("lor.index")
local cookie_middleware = require("lor.lib.middleware.cookie")
local check_login_middleware = require("app.middleware.check_login")
local whitelist = require("app.config.config").whitelist
local router = require("app.router")

-- 全局 app 对象, Golang 通过调用app.run来处理每个请求
app = lor()

app:conf("view enable", true)
-- app:conf("view engine", "tmpl")
app:conf("view ext", "html")
app:conf("views", "app/views")

app:use(cookie_middleware())

-- filter: add response header
app:use(function(req, res, next)
    res:set_header('X-Powered-By', 'Lor Framework')
    next()
end)

-- intercepter: login or not
app:use(check_login_middleware(whitelist))

router(app) -- business routers and routes

-- 404 error
-- app:use(function(req, res, next)
--     if req:is_found() ~= true then
--         res:status(404):send("404! sorry, page not found.")
--     end
-- end)

-- error handle middleware
app:erroruse(function(err, req, res, next)
    print(err)
    if req:is_found() ~= true then
         res:status(404):send("404! sorry, page not found. uri:" .. req.path)
    else
        res:status(500):send("unknown error")
    end
end)


--ngx.say(app.router.trie:gen_graph())

-- app:run()
