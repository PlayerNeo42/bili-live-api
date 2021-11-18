// 得到账号登录信息
console.log(JSON.stringify(document.cookie.split(';').reduce((prev, current) => {
    const [k, v] = current.split('=').map(v => v.trim());
    return (k === 'SESSDATA' || k === 'bili_jct' || k === 'buvid3') ? {...prev, [k]: v}: prev ;
}, {"SESSDATA":""}), null, 2))
