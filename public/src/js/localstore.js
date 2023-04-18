const keyPrefix = "save_i_"

const getKey = (key) => {
    return keyPrefix + key
}

export const lsave = (key, value) => {
    console.log(key, value)
    let valueText = JSON.stringify(value)
    localStorage.setItem(getKey(key), valueText)
}


// Added for storing successfully connected sessions
export const ladd = (key, value) => {
    value.value = value.server_url
    var currentValue = localStorage.getItem(getKey(key))
    if (currentValue === null) {
        lsave(key, [value])
        return
    }
    currentValue = JSON.parse(currentValue)
    if (Array.isArray(currentValue)) {
        for (let i = 0; i < currentValue.length; i++) {
            if (currentValue[i].value === value.value) {
                return
            }
        }
        currentValue.push(value)
        lsave(key, currentValue)
    }
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
