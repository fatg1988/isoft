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
    companysite:"^http[s]?:\\/\\/([\\w-]+\\.)+[\\w-]+([\\w-./?%&#=]*)?$",
    normal:"^[A-Za-z0-9_\\-\\u4e00-\\u9fa5]+$",  // 正常字符,中间不含空格
    normalWithEmpty:"^[A-Za-z0-9_\\-\\s\\u4e00-\\u9fa5]+$"  // 正常字符,可含有空格
};

//主函数
// 主要分为两大类,一类是使用 validateSettings 中的默认回调函数(onFocus\isNull\error\succeed)
// 另一类是使用 ilearningValidate 中传入的回调函数
(function ($) {
    $.fn.ilearningValidate = function (option, callback, def) {
        var ele = this;
        var id = ele.attr("id");                                                // input 元素的 id 属性值
        var type = ele.attr("type");                                            // input 元素的 type 属性
        var rel = ele.attr("rel");                                              // 用于标记当前元素是否是 select 或者 textarea
        var _onFocus = $("#" + id + validateSettings.onFocus.container);        // focus 情况下回显信息的 label
        var _succeed = $("#" + id + validateSettings.succeed.container);        // succeed 情况下回显信息的 label
        var _isNull = $("#" + id + validateSettings.isNull.container);          // isNull 情况下回显信息的 label
        var _error = $("#" + id + validateSettings.error.container);            // error 情况下回显信息的 label
        if (def == true) {
            var str = ele.val();
            var tag = ele.attr("sta");
            if (str == "" || str == "-1" || str == "请选择") {             // 表单元素 value 值为空串或者 -1 (select 请选择 value 为 -1) 时
                validateSettings.isNull.run({
                    prompts:option,
                    element:ele,
                    isNullEle:_isNull,
                    succeedEle:_succeed
                }, option.isNull);
            } else if (tag == 1 || tag == 2) {
                return;
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
                case "text":            // 表单元素 type 为 text 或者 password 时绑定 focus 和 blur 事件
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
                    } else if(rel && rel == "textarea"){
                        ele.bind("focus", function () {
                            var str = ele.val();
                            if (str == def || $.trim(str) == "") {
                                ele.val("");        // 默认值置空处理
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
    },
    isNormal:function (str) {
        return new RegExp(validateRegExp.normal).test(str);
    },
    normalWithEmpty:function (str) {
        return new RegExp(validateRegExp.normalWithEmpty).test(str);
    }
};

// 验证文本
var validatePrompt = {

};

// 回调函数
var validateFunction = {
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

