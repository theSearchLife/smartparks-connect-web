const keyPrefix = "save_i_"

const getKey = (key) => {
    return keyPrefix + key
}

export const lsave = (key, value) => {
    console.log(key, value)
    let valueText = JSON.stringify(value)
    localStorage.setItem(getKey(key), valueText)
}

export const lget = (key) => {
    let valueText = localStorage.getItem(getKey(key))
    if (valueText == "") {
        return valueText
    }
    return JSON.parse(valueText)
}

export const getSet = (key, obj) => {
    let value = lget(key)
    if (value != null) {
        for (var i in obj) {
            obj[i] = value[i]
        }
        for (var i in value) {
            obj[i] = value[i]
        }
    }
    return obj
}
