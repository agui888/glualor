local lor = require("lor.index")
app = lor()

app:conf("view enable", true)

app:get("/", function(req, res, next)
    res:send("hello world!")
end)

app:get("/json", function(req, res, next)
    res:json(req.query)
end)

app:get("/redirect", function(req, res, next)
    res:redirect("http://www.qq.com")
end)

app:get("/template", function(req, res, next)
    res:render("template/index", {name = req.query.name or "Tim"})
end)
