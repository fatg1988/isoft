// 验证正则表达式汇总
var validateRegExp = {
    decmal:"^([+-]?)\\d*\\.\\d+$", //浮点数
    decmal1:"^[1-9]\\d*.\\d*|0.\\d*[1-9]\\d*$", //正浮点数
    decmal2:"^-([1-9]\\d*.\\d*|0.\\d*[1-9]\\d*)$", //负浮点数
    decmal3:"^-?([1-9]\\d*.\\d*|0.\\d*[1-9]\\d*|0?.0+|0)$", //浮点数
    decmal4:"^[1-9]\\d*.\\d*|0.\\d*[1-9]\\d*|0?.0+|0$", //非负浮点数（正浮点数 + 0）
    decmal5:"^(-([1-9]\\d*.\\d*|0.\\d*[1-9]\\d*))|0?.0+|0$", //非正浮点数（负浮点数 + 0）
    intege:"^-?[1-9]\\d*$", //整数
    intege1:"^[1-9]\\d*$", //正整数
    intege2:"^-[1-9]\\d*$", //负整数
    num:"^([+-]?)\\d*\\.?\\d+$", //数字
    num1:"^[1-9]\\d*|0$", //正数（正整数 + 0）
    num2:"^-[1-9]\\d*|0$", //负数（负整数 + 0）
    ascii:"^[\\x00-\\xFF]+$", //仅ACSII字符
    chinese:"^[\\u4e00-\\u9fa5]+$", //仅中文
    color:"^[a-fA-F0-9]{6}$", //颜色
    date:"^\\d{4}(\\-|\\/|\.)\\d{1,2}\\1\\d{1,2}$", //日期
    email:"^\\w+((-\\w+)|(\\.\\w+))*\\@[A-Za-z0-9]+((\\.|-)[A-Za-z0-9]+)*\\.[A-Za-z0-9]+$", //邮件
    idcard:"^[1-9]([0-9]{14}|[0-9]{17})$", //身份证
    ip4:"^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$", //ip地址
    letter:"^[A-Za-z]+$", //字母
    letter_l:"^[a-z]+$", //小写字母
    letter_u:"^[A-Z]+$", //大写字母
    mobile:"^0?(13|15|18)[0-9]{9}$", //手机
    notempty:"^\\S+$", //非空
    password:"^.*[A-Za-z0-9\\w_-]+.*$", //密码
    fullNumber:"^[0-9]+$", //数字
    picture:"(.*)\\.(jpg|bmp|gif|ico|pcx|jpeg|tif|png|raw|tga)$", //图片
    qq:"^[1-9]*[1-9][0-9]*$", //QQ号码
    rar:"(.*)\\.(rar|zip|7zip|tgz)$", //压缩文件
    tel:"^[0-9\-()（）]{7,18}$", //电话号码的函数(包括验证国内区号,国际区号,分机号)
    url:"^http[s]?:\\/\\/([\\w-]+\\.)+[\\w-]+([\\w-./?%&=]*)?$", //url
    username:"^[A-Za-z0-9_\\-\\u4e00-\\u9fa5]+$", //用户名
    deptname:"^[A-Za-z0-9_()（）\\-\\u4e00-\\u9fa5]+$", //单位名
    zipcode:"^\\d{6}$", //邮编
    realname:"^[A-Za-z\\u4e00-\\u9fa5]+$", // 真实姓名
    companyname:"^[A-Za-z0-9_()（）\\-\\u4e00-\\u9fa5]+$",
    companyaddr:"^[A-Za-z0-9_()（）\\#\\-\\u4e00-\\u9fa5]+$",
    companysite:"^http[s]?:\\/\\/([\\w-]+\\.)+[\\w-]+([\\w-./?%&#=]*)?$"
};

//主函数
// 主要分为两大类,一类是使用 validateSettings 中的默认回调函数(onFocus\isNull\error\succeed)
// 另一类是使用 ilearningValidate 中传入的回调函数
(function ($) {
    $.fn.ilearningValidate = function (option, callback, def) {
        var ele = this;
        var id = ele.attr("id");                                                // input 元素的 id 属性值
        var type = ele.attr("type");                                            // input 元素的 type 属性
        var rel = ele.attr("rel");                                              // 用于标记当前元素是否是 select
        var _onFocus = $("#" + id + validateSettings.onFocus.container);        // focus 情况下回显信息的 label
        var _succeed = $("#" + id + validateSettings.succeed.container);        // succeed 情况下回显信息的 label
        var _isNull = $("#" + id + validateSettings.isNull.container);          // isNull 情况下回显信息的 label
        var _error = $("#" + id + validateSettings.error.container);            // error 情况下回显信息的 label
        if (def == true) {
            var str = ele.val();
            var tag = ele.attr("sta");
            if (str == "" || str == "-1") {             // 表单元素 value 值为空串或者 -1 时
                validateSettings.isNull.run({
                    prompts:option,
                    element:ele,
                    isNullEle:_isNull,
                    succeedEle:_succeed
                }, option.isNull);
            } else if (tag == 1 || tag == 2) {

            } else {                                    // 非空时执行
                // 回调函数 5 大参数 option = {prompts:option,element:ele,value:str,errorEle:_error,succeedEle:_succeed}
                callback({
                    prompts:option,
                    element:ele,
                    value:str,
                    errorEle:_error,                    // error 情况下回显信息的 label
                    succeedEle:_succeed                 // succeed 情况下回显信息的 label
                });
            }
        } else {
            if (typeof def == "string") {               // 默认值 def 是字符串类型时
                ele.val(def);                           // 给元素设置默认值 def
            }
            if (type == "checkbox" || type == "radio") {    // 复选框或者单选框
                if (ele.attr("checked") == true) {          // 选中时
                    ele.attr("sta", validateSettings.succeed.state);    // 给表单元素设置标记 sta 为 succeed 已执行状态
                }
            }
            switch (type) {
                case "text":
                case "password":
                    ele.bind("focus", function () {
                        var str = ele.val();
                        if (str == def) {
                            ele.val("");        // password 类型 focus 时如果是默认值则置空
                        }
                        if (id == "pwd") {
                            $("#pwdstrength").hide();   // 隐藏密码强度 label 条
                        }
                        validateSettings.onFocus.run({  // 触发 focus 时的回调函数
                            prompts:option,
                            element:ele,
                            value:str,
                            onFocusEle:_onFocus,
                            succeedEle:_succeed
                        }, option.onFocus);
                    })
                        .bind("blur", function () {
                            var str = ele.val();        // 失去焦点时如果为空串则使用默认值
                            if (str == "") {
                                ele.val(def);
                            }
                            if (validateRules.isNull(str)) {
                                validateSettings.isNull.run({   // 触发为空时的回调函数
                                    prompts:option,
                                    element:ele,
                                    value:str,
                                    isNullEle:_isNull,
                                    succeedEle:_succeed
                                }, "");
                            } else {
                                callback({              // 非空时执行回调函数(并传入回调函数 6 大参数)
                                    prompts:option,
                                    element:ele,
                                    value:str,
                                    errorEle:_error,
                                    isNullEle:_isNull,
                                    succeedEle:_succeed
                                });
                            }
                        });
                    break;
                default:
                    if (rel && rel == "select") {                   // 是 select 类型的
                        ele.bind("change", function () {            // 给元素绑定 change 事件
                            var str = ele.val();                    // 执行回调函数(并传入回调函数 6 大参数)
                            callback({
                                prompts:option,
                                element:ele,
                                value:str,
                                errorEle:_error,
                                isNullEle:_isNull,
                                succeedEle:_succeed
                            });
                        })
                    } else {
                        ele.bind("click", function () {            // 不满足以上条件的 default 类型
                            callback({
                                prompts:option,
                                element:ele,
                                errorEle:_error,
                                isNullEle:_isNull,
                                succeedEle:_succeed
                            });
                        })
                    }
                    break;
            }
        }
    }
})(jQuery);

//配置,主要是 onFocus\isNull\error\succeed 几种默认场景的回调函数(处理函数)
var validateSettings = {
    onFocus:{
        state:null,
        container:"_error",
        style:"focus",
        run:function (option, str) {
            if (!validateRules.checkType(option.element)) {  // 不是 checkbox\radio\select 中的任意一种
                // 切换表单元素的高亮显示状态
                option.element.removeClass(validateSettings.INPUT_style2).addClass(validateSettings.INPUT_style1);
            }
            // focus 情况下回显提示信息的 label 移除 focus 样式
            option.onFocusEle.removeClass().addClass(validateSettings.onFocus.style).html(str);
        }
    },
    isNull:{
        state:0,
        container:"_error",
        style:"null",
        run:function (option, str) {
            option.element.attr("sta", 0);          // 为空时设置标记状态为 0,用于禁止提交
            if (!validateRules.checkType(option.element)) { // 不是 checkbox\radio\select 中的任意一种
                if (str != "") {
                    option.element.removeClass(validateSettings.INPUT_style1).addClass(validateSettings.INPUT_style2);
                } else {
                    option.element.removeClass(validateSettings.INPUT_style2).removeClass(validateSettings.INPUT_style1);
                }
            }
            option.succeedEle.removeClass(validateSettings.succeed.style);
            option.isNullEle.removeClass().addClass(validateSettings.isNull.style).html(str);
        }
    },
    error:{
        state:1,
        container:"_error",
        style:"error",
        run:function (option, str) {
            option.element.attr("sta", 1);
            if (!validateRules.checkType(option.element)) {
                option.element.removeClass(validateSettings.INPUT_style1).addClass(validateSettings.INPUT_style2);
            }
            option.succeedEle.removeClass(validateSettings.succeed.style);
            option.errorEle.removeClass().addClass(validateSettings.error.style).html(str);
        }
    },
    succeed:{
        state:2,
        container:"_succeed",
        style:"succeed",
        run:function (option) {
            option.element.attr("sta", 2);
            option.errorEle.empty();
            if (!validateRules.checkType(option.element)) {
                option.element.removeClass(validateSettings.INPUT_style1).removeClass(validateSettings.INPUT_style2);
            }
            if (option.element.attr("id") == "schoolinput" && $("#schoolid").val() == "") {
                return;
            }
            option.succeedEle.addClass(validateSettings.succeed.style);
        }
    },
    INPUT_style1:"highlight1",
    INPUT_style2:"highlight2"

};

//验证规则
var validateRules = {
    isNull:function (str) {
        return (str == "" || typeof str != "string");
    },
    betweenLength:function (str, _min, _max) {
        return (str.length >= _min && str.length <= _max);
    },
    isUid:function (str) {
        return new RegExp(validateRegExp.username).test(str);
    },
    fullNumberName:function (str) {
        return new RegExp(validateRegExp.fullNumber).test(str);
    },
    isPwd:function (str) {
        return /^.*([\W_a-zA-z0-9-])+.*$/i.test(str);
    },
    isPwd2:function (str1, str2) {
        return (str1 == str2);
    },
    isEmail:function (str) {
        return new RegExp(validateRegExp.email).test(str);
    },
    isTel:function (str) {
        return new RegExp(validateRegExp.tel).test(str);
    },
    isMobile:function (str) {
        return new RegExp(validateRegExp.mobile).test(str);
    },
    checkType:function (element) {
        return (element.attr("type") == "checkbox" || element.attr("type") == "radio" || element.attr("rel") == "select");
    },
    isChinese:function (str) {
        return new RegExp(validateRegExp.chinese).test(str);
    },
    isRealName:function (str) {
        return new RegExp(validateRegExp.realname).test(str);
    },
    isDeptname:function (str) {
        return new RegExp(validateRegExp.deptname).test(str);
    },
    isCompanyname:function (str) {
        return new RegExp(validateRegExp.companyname).test(str);
    },
    isCompanyaddr:function (str) {
        return new RegExp(validateRegExp.companyaddr).test(str);
    },
    isCompanysite:function (str) {
        return new RegExp(validateRegExp.companysite).test(str);
    }
};

// 验证文本
var validatePrompt = {
    username:{
        onFocus:"4-20位字符，可由中文、英文、数字及“_”、“-”组成",
        succeed:"",
        isNull:"请输入用户名",
        error:{
            beUsed:"该用户名已被使用，请使用其它用户名注册，如果您是&quot;{1}&quot;，请<a href='https://passport.360buy.com/new/login.aspx' class='flk13'>登录</a>",
            badLength:"用户名长度只能在4-20位字符之间",
            badFormat:"用户名只能由中文、英文、数字及“_”、“-”组成",
            fullNumberName:"用户名不能全为数字"
        }
    },
    pwd:{
        onFocus:"6-16位字符，可由英文、数字及标点符号组成",
        succeed:"",
        isNull:"请输入密码",
        error:{
            badLength:"密码长度只能在6-16位字符之间",
            badFormat:"密码只能由英文、数字及标点符号组成"
        }
    },
    pwd2:{
        onFocus:"请再次输入密码",
        succeed:"",
        isNull:"请输入密码",
        error:{
            badLength:"密码长度只能在6-16位字符之间",
            badFormat2:"两次输入密码不一致",
            badFormat1:"密码只能由英文、数字及标点符号组成"
        }
    },
    mail:{
        onFocus:"请输入常用的邮箱，将用来找回密码、接收订单通知等",
        succeed:"",
        isNull:"请输入邮箱",
        error:{
            beUsed:"该邮箱已被使用，请更换其它邮箱，或使用该邮箱<a href='http://passport.360buy.com/retrievepassword.aspx' class='flk13'>找回密码</a>",
            badFormat:"邮箱格式不正确",
            badLength:"您填写的邮箱过长，邮件地址只能在50个字符以内"
        }
    },
    authcode:{
        onFocus:"请输入图片中的字符，不区分大小写",
        succeed:"",
        isNull:"请输入验证码",
        error:"验证码错误"
    },
    protocol:{
        onFocus:"",
        succeed:"",
        isNull:"请先阅读并同意《京东商城用户协议》",
        error:""
    },
    referrer:{
        onFocus:"如您注册并完成订单，推荐人有机会获得积分",
        succeed:"",
        isNull:"",
        error:""
    },
    schoolinput:{
        onFocus:"您可以用简拼、全拼、中文进行校名模糊查找",
        succeed:"",
        isNull:"请填选学校名称",
        error:"请填选学校名称"
    },
    empty:{
        onFocus:"",
        succeed:"",
        isNull:"",
        error:""
    }
};

var nameold, emailold, authcodeold;
var namestate = false, emailstate = false, authcodestate = false;
// 回调函数
var validateFunction = {
    // 回调函数 5 大参数 option = {prompts:option,element:ele,value:str,errorEle:_error,succeedEle:_succeed}
    username:function (option) {

        var format = validateRules.isUid(option.value);     // 是否满足用户名格式
        var length = validateRules.betweenLength(option.value.replace(/[^\x00-\xff]/g, "**"), 4, 20);       // 长度是否符合
        if (!length && format) {            // 格式正确但长度不符合
            validateSettings.error.run(option, option.prompts.error.badLength);
        }
        else if (!length && !format) {      // 长度和格式都不符合
            validateSettings.error.run(option, option.prompts.error.badFormat);
        }
        else if (length && !format) {       // 长度符合但格式不符合
            validateSettings.error.run(option, option.prompts.error.badFormat);
        } else if (validateRules.fullNumberName(option.value)) {    // 全数字
            validateSettings.error.run(option, option.prompts.error.fullNumberName);
        } else {
            if (!namestate || nameold != option.value) {
                if (nameold != option.value) {              // 不相等,代表第一次检查,需要后台验证
                    nameold = option.value;
                    option.errorEle.html("<span style='color:#999'>检验中……</span>");  // 验证用户名
                    $.getJSON("AjaxService.aspx?action=CheckUnicknme&uid=" + escape(option.value) + "&r=" + Math.random(), function (date) {
                        if (date.success == 0) {            // 用户名验证成功
                            validateSettings.succeed.run(option);
                            namestate = true;
                        } else {                            // 用户名验证失败
                            validateSettings.error.run(option, option.prompts.error.beUsed.replace("{1}", option.value));
                            namestate = false;
                        }
                    })
                }
                else {                                  // 相等则代表已经存在过,直接提示异常信息
                    validateSettings.error.run(option, option.prompts.error.beUsed.replace("{1}", option.value));
                    namestate = false;
                }
            }
            else {      // !(!namestate || nameold != option.value) 代表已经验证通过,执行 succeed 的回调函数
                validateSettings.succeed.run(option);
            }
        }
    },
    pwd:function (option) {
        var str1 = option.value;
        var str2 = $("#pwd2").val();
        var format = validateRules.isPwd(option.value);
        var length = validateRules.betweenLength(option.value, 6, 16);
        $("#pwdstrength").hide();
        if (!length && format) {
            validateSettings.error.run(option, option.prompts.error.badLength);
        }
        else if (!length && !format) {
            validateSettings.error.run(option, option.prompts.error.badFormat);
        }
        else if (length && !format) {
            validateSettings.error.run(option, option.prompts.error.badFormat);
        }
        else {
            validateSettings.succeed.run(option);
            validateFunction.pwdstrength();
        }
        if (str2 == str1) {
            $("#pwd2").focus();
        }
    },
    pwd2:function (option) {
        var str1 = option.value;
        var str2 = $("#pwd").val();
        var length = validateRules.betweenLength(option.value, 6, 16);
        var format2 = validateRules.isPwd2(str1, str2);
        var format1 = validateRules.isPwd(str1);
        if (!length) {
            validateSettings.error.run(option, option.prompts.error.badLength);
        } else {
            if (!format1) {
                validateSettings.error.run(option, option.prompts.error.badFormat1);
            } else {
                if (!format2) {
                    validateSettings.error.run(option, option.prompts.error.badFormat2);
                }
                else {
                    validateSettings.succeed.run(option);
                }
            }
        }
    },
    mail:function (option) {
        var format = validateRules.isEmail(option.value);
        var format2 = validateRules.betweenLength(option.value, 0, 50);
        if (!format) {          // 邮箱格式不正确
            validateSettings.error.run(option, option.prompts.error.badFormat);
        } else {
            if (!format2) {     // 邮箱长度不正确
                validateSettings.error.run(option, option.prompts.error.badLength);
            } else {
                if (!emailstate || emailold != option.value) {  // 验证状态失败需要重新验证,旧值与新值不一样也需要重新验证
                    if (emailold != option.value) {
                        emailold = option.value;
                        option.errorEle.html("<span style='color:#999'>检验中……</span>");
                        $.getJSON("AjaxService.aspx?action=CheckUemail&str=" + escape(option.value) + "&r=" + Math.random(), function (date) {
                            if (date.success == 0) {        // 验证成功
                                validateSettings.succeed.run(option);
                                emailstate = true;          // 标记验证状态
                            } else {                        // 验证失败
                                validateSettings.error.run(option, option.prompts.error.beUsed);
                                emailstate = false;
                            }
                        })
                    }
                    else {  // 相等则代表已经存在过,直接提示异常信息
                        validateSettings.error.run(option, option.prompts.error.beUsed);
                        emailstate = false;
                    }
                }
                else {  // !(!emailstate || emailold != option.value) 代表已经验证通过,执行 succeed 的回调函数
                    validateSettings.succeed.run(option);
                }
            }
        }
    },
    referrer:function (option) {
        var bool = validateRules.isNull(option.value);
        if (bool) {
            option.element.val("可不填");

        } else {
            validateSettings.succeed.run(option);
        }
    },
    schoolinput:function (option) {
        var bool = validateRules.isNull(option.value);
        if (bool) {
            validateSettings.error.run(option, option.prompts.error);

        } else {
            validateSettings.succeed.run(option);
        }
    },
    authcode:function (option) {
//        if (!authcodestate || authcodeold != option.value) {
//            if (authcodeold != option.value) {
////                authcodeold = option.value;
////                option.errorEle.html("<span style='color:#999'>检验中……</span>");
////                var uuid = $("#JD_Verification1").attr("src").split("&uid=")[1].split("&")[0];
////                $.getJSON("AjaxService.aspx?action=CheckAuthcode&str=" + escape(option.value) + "&r=" + Math.random() + "&uuid=" + uuid, function(date) {
////                    if (date.success == 0) {
////                        validateSettings.succeed.run(option);
////                        authcodestate = true;
////                    } else {
////                        validateSettings.error.run(option, option.prompts.error);
////                        authcodestate = false;
////                    }
////                })
//            }
//            else {
//                validateSettings.error.run(option, option.prompts.error);
//                authcodestate = false;
//            }
//        }
//        else {
//            validateSettings.succeed.run(option);
//        }

        validateSettings.succeed.run(option);
        authcodestate = true;
    },
    protocol:function (option) {
        if (option.element.attr("checked") == true) {
            option.element.attr("sta", validateSettings.succeed.state);
            option.errorEle.html("");
        } else {
            option.element.attr("sta", validateSettings.isNull.state);
            option.succeedEle.removeClass(validateSettings.succeed.style);
        }
    },
    pwdstrength:function () {

        var element = $("#pwdstrength");
        var value = $("#pwd").val();
        if (value.length >= 6 && validateRules.isPwd(value)) {
            $("#pwd_error").empty();
            element.show();

            var pattern_1 = /^.*([\W_])+.*$/i;
            var pattern_2 = /^.*([a-zA-Z])+.*$/i;
            var pattern_3 = /^.*([0-9])+.*$/i;
            var level = 0;

            if (value.length > 10) {
                level++;
            }

            if (pattern_1.test(value)) {
                level++;
            }

            if (pattern_2.test(value)) {
                level++;
            }

            if (pattern_3.test(value)) {
                level++;
            }

            if (level > 3) {
                level = 3;
            }

            switch (level) {
                case 1:
                    element.removeClass().addClass("strengthA");
                    break;
                case 2:
                    element.removeClass().addClass("strengthB");
                    break;
                case 3:
                    element.removeClass().addClass("strengthC");
                    break;
                default:
                    break;
            }
        } else {
            element.hide();
        }
    },
    checkGroup:function (elements) {
        for (var i = 0; i < elements.length; i++) {
            if (elements[i].checked) {
                return true;
            }
        }
        return false;
    },
    checkSelectGroup:function (elements) {
        for (var i = 0; i < elements.length; i++) {
            if (elements[i].value == -1) {
                return false;
            }
        }
        return true;
    },
    showPassword:function (type) {
        var v1 = $("#pwd").val(), s1 = $("#pwd").attr("sta"), c1 = document.getElementById("pwd").className, t1 = $("#pwd").attr("tabindex");
        var v2 = $("#pwd2").val(), s2 = $("#pwd2").attr("sta"), c2 = document.getElementById("pwd2").className, t2 = $("#pwd2").attr("tabindex");
        var P1 = $("<input type='" + type + "' value='" + v1 + "' sta='" + s1 + "' class='" + c1 + "' id='pwd' name='pwd' tabindex='" + t1 + "'/>");
        $("#pwd").after(P1).remove();
        $("#pwd").bind("keyup",
            function () {
                validateFunction.pwdstrength();
            }).ilearningValidate(validatePrompt.pwd, validateFunction.pwd);
        var P2 = $("<input type='" + type + "' value='" + v2 + "' sta='" + s2 + "' class='" + c2 + "' id='pwd2' name='pwd2' tabindex='" + t2 + "'/>");
        $("#pwd2").after(P2).remove();
        $("#pwd2").ilearningValidate(validatePrompt.pwd2, validateFunction.pwd2);
    },
    FORM_submit:function (elements) {
        var bool = true;
        for (var i = 0; i < elements.length; i++) {
            if ($(elements[i]).attr("sta") == 2) {
                bool = true;
            } else {
                bool = false;
                break;
            }
        }
        return bool;
    }
};

