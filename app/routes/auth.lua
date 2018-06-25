local pairs = pairs
local ipairs = ipairs
local lor = require("lor.index")
local users = require("app.config.config").users
local authRouter = lor:Router()

authRouter:get("/login", function(req, res, next)
    res:render("login")
end)

authRouter:post("/login", function(req, res, next)
    local username = req.body.username 
    local password = req.body.password

    local isExist = false
    for i, v in ipairs(users) do
        if v.username == username and v.password == password then
            req.cookie.set({Name = "username", Value = username, Path="/"})
            isExist = true
            return res:redirect("/todo/index")
        end
    end

    if not isExist then
        res:redirect("/error/",{
            errMsg = "Wrong username or password! Please check."
        })
    end
end)

authRouter:get("/logout", function(req, res, next)
    req.cookie.set({Name = "username", Value = ""})
    res:redirect("/auth/login")
end)

return authRouter
