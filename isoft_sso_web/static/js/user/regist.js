function regist() {
    var username = $("input[name='username']").val();
    var passwd = $("input[name='passwd']").val();
    var proxy = $("input[name='proxy']:checked").val();
    if (proxy != "proxy") {
        $("._proxy_error").show();
        return;
    } else {
        $("._proxy_error").hide();
    }

    if (passwd.length < 10) {
        $("._password_error").show();
        return;
    } else {
        $("._password_error").hide();
    }

    $.ajax({
        url: "/user/regist/",
        method: "post",
        data: {"username": username, "passwd": passwd, "proxy": proxy},
        success: function (data) {
            var result = JSON.parse(data);
            if(result.Status == "SUCCESS"){
                window.location.href="/index"
            }else{
                if(result.ErrorCode == "用户已注册!"){
                    $("._username_error").show();
                }else{
                    window.location.href="/user/login/"
                }
            }
        }
    })
}