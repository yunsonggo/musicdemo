import originJSONP from 'jsonp'
// https://github.com/webmodules/jsonp 
// npm install jsonp 

export default function jsonp(url = '', data = {}, option) {
    return new Promise((resolve, reject) => {
        let dataStr = ''
        Object.keys(data).forEach(key => {
            dataStr += key + '=' + data[key] + '&'
        })
        if (dataStr !== '') {
            dataStr = dataStr.substring(0, dataStr.lastIndexOf('&'))
            url = url + '?' + dataStr
        }
        originJSONP(url, option, (err, data) => {
            if (!err) {
                resolve(data)
            } else {
                reject(err)
            }
        })
    })
}

