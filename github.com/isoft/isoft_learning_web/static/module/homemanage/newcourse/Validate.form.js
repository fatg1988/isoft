// 扩展验证文本
$.extend(validatePrompt, {
    course_name:{
        onFocus:"请不要使用特殊字符",
        succeed:"",
        isNull:"请输入课程名称",
        error:{
            badLength:"课程名称1-20个字符",
            badFormat:"课程名称格式不正确,含有非法字符",
            beUsed:"该课程名称已经使用,请重新输入"
        }
    },
    course_type:{
        onFocus:"",
        succeed:"",
        isNull:"课程类型不能为空",
        error:""
    },
    course_sub_type:{
        onFocus:"请不要使用特殊字符",
        succeed:"",
        isNull:"请输入课程子类别名称",
        error:{
            badLength:"课程子类型名称1-20个字符",
            badFormat:"课程子类型名称格式不正确,含有非法字符"
        }
    },
    course_short_desc:{
        onFocus:"请不要使用特殊字符",
        succeed:"",
        isNull:"请输入课程简介信息",
        error:{
            badLength:"请输入课程简介信息1-50个字符",
            badFormat:"课程简介信息格式不正确,含有非法字符"
        }
    }
});

var course_name_old;
var course_name_state = false;

// 扩展回调函数
$.extend(validateFunction, {
    // 回调函数 5 大参数 option = {prompts:option,element:ele,value:str,errorEle:_error,succeedEle:_succeed}
    course_name:function (option) {
        var length = validateRules.betweenLength($.trim(option.value), 1, 20);
        var format = validateRules.isNormal($.trim(option.value));
        if (!length) {
            validateSettings.error.run(option, option.prompts.error.badLength);
        } else if (!format) {
            validateSettings.error.run(option, option.prompts.error.badFormat);
        } else {
            if (!course_name_state || course_name_old != option.value) {
                if (course_name_old != option.value) {              // 不相等,代表第一次检查,需要后台验证
                    course_name_old = option.value;
                    option.errorEle.html("<span style='color:#999'>检验中……</span>");  // 验证用户名
                    $.getJSON("/course/queryCourseExist?course_name=" + option.value + "&r=" + Math.random(), function (data) {
                        var obj = $.parseJSON(data);
                        if (obj.flag == false) {            // course_name 不存在
                            validateSettings.succeed.run(option);
                            course_name_state = true;
                        } else {                            // course_name 已经存在
                            validateSettings.error.run(option, option.prompts.error.beUsed.replace("{1}", option.value));
                            course_name_state = false;
                        }
                    })
                }
                else {                                  // 相等则代表已经存在过,直接提示异常信息
                    validateSettings.error.run(option, option.prompts.error.beUsed.replace("{1}", option.value));
                    course_name_state = false;
                }
            }
            else {      // 代表已经验证通过,执行 succeed 的回调函数
                validateSettings.succeed.run(option);
            }
        }
    },
    course_type:function(option) {
        var bool = (option.value == -1 || option.value == "请选择");
        if (bool) {
            validateSettings.isNull.run(option, "");
        }
        else {
            validateSettings.succeed.run(option);
        }
    },
    course_sub_type:function (option) {
        var length = validateRules.betweenLength($.trim(option.value), 1, 20);
        var format = validateRules.isNormal($.trim(option.value));
        if (!length) {
            validateSettings.error.run(option, option.prompts.error.badLength);
        } else if (!format) {
            validateSettings.error.run(option, option.prompts.error.badFormat);
        } else {
            validateSettings.succeed.run(option);
        }
    },
    course_short_desc:function (option) {
        var length = validateRules.betweenLength($.trim(option.value), 1, 50);
        var format = validateRules.normalWithEmpty($.trim(option.value));
        if (!length) {
            validateSettings.error.run(option, option.prompts.error.badLength);
        } else if (!format) {
            validateSettings.error.run(option, option.prompts.error.badFormat);
        } else {
            validateSettings.succeed.run(option);
        }
    },
    FORM_validate:function() {
        $("#course_name").ilearningValidate(validatePrompt.course_name, validateFunction.course_name, true);
        $("#course_type").ilearningValidate(validatePrompt.course_type, validateFunction.course_type, true);
        $("#course_sub_type").ilearningValidate(validatePrompt.course_sub_type, validateFunction.course_sub_type, true);
        $("#course_short_desc").ilearningValidate(validatePrompt.course_short_desc, validateFunction.course_short_desc, true);
        return validateFunction.FORM_submit(["#course_name","#course_sub_type","#course_short_desc","#course_type"]);
    }
});

$(function () {
    //默认下 course_name 框获得焦点
    setTimeout(function() {
        $("#course_name").get(0).focus();
    }, 0);

    // course_name 验证
    $("#course_name").ilearningValidate(validatePrompt.course_name, validateFunction.course_name);
    // course_type 验证
    $("#course_type").ilearningValidate(validatePrompt.course_type, validateFunction.course_type);
    // course_sub_type 验证
    $("#course_sub_type").ilearningValidate(validatePrompt.course_sub_type, validateFunction.course_sub_type);
    // sourse_short_desc 验证
    $("#course_short_desc").ilearningValidate(validatePrompt.course_short_desc, validateFunction.course_short_desc);

    //表单提交验证和服务器请求
    $("#registsubmit").click(function() {
        var flag = validateFunction.FORM_validate();
        if (flag) {
            $(this).attr({"disabled":"disabled"}).attr({"value":"提交中,请稍等"});
            $.ajax({
                type: "POST",
                url: "/course/newcourse/add",
                data: $("#formpersonal").serialize(),
                success: function(data) {
                    var obj = JSON.parse(data);
                    if (obj.status == "SUCCESS") {
                        window.location = "/course/courselist";
                    }else{
                        $("#course_short_desc_error").removeClass().addClass("error").html(obj.msg);
                    }
                }
            });
        }
    });
});

