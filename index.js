const http = require('http')

const fetch = require('node-fetch')
const cheerio = require('cheerio')

const server = http.createServer(async (req, res) => {
  const defcon_res = await fetch('https://www.defconlevel.com/current-level.php')
  const $ = cheerio.load(await defcon_res.text())

  const defcon_level = /:.+(\d)/.exec($('.header-defcon-level').text())[1]

  res.writeHead(200, { 'Content-Type': 'text/plain' })

  res.write(defcon_level)

  res.end()
});

const port = process.env.PORT || 5000
server.listen(port, () => {
  console.log(`listening on ${port}`)
})