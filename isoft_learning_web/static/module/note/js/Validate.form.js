// 扩展验证文本
$.extend(validatePrompt, {
    note_name:{
        onFocus:"请不要使用特殊字符",
        succeed:"",
        isNull:"请输入笔记名称",
        error:{
            badLength:"笔记名称1-20个字符",
            badFormat:"笔记名称格式不正确,含有非法字符",
            beUsed:"该笔记名称已经使用,请重新输入"
        }
    },
    note_key_words:{
        onFocus:"请不要使用特殊字符",
        succeed:"",
        isNull:"请输入检索关键词",
        error:{
            badLength:"检索关键词1-50个字符",
            badFormat:"检索关键词格式不正确,含有非法字符"
        }
    }
});

var note_name_old;
var note_name_state = false;

// 扩展回调函数
$.extend(validateFunction, {
    // 回调函数 5 大参数 option = {prompts:option,element:ele,value:str,errorEle:_error,succeedEle:_succeed}
    note_name:function (option) {
        var length = validateRules.betweenLength($.trim(option.value), 1, 20);
        var format = validateRules.isNormal($.trim(option.value));
        if (!length) {
            validateSettings.error.run(option, option.prompts.error.badLength);
        } else if (!format) {
            validateSettings.error.run(option, option.prompts.error.badFormat);
        } else {
            // 不是 isFormReload 表单回显需要验证用户名是否重复
            if ((!note_name_state || note_name_old != option.value) && !isFormReload()) {
                if (note_name_old != option.value) {              // 不相等,代表第一次检查,需要后台验证
                    note_name_old = option.value;
                    option.errorEle.html("<span style='color:#999'>检验中……</span>");  // 验证用户名
                    $.getJSON("/note/queryNoteExist?note_name=" + option.value + "&r=" + Math.random(), function (data) {
                        if (data.flag == false) {            // course_name 不存在
                            validateSettings.succeed.run(option);
                            note_name_state = true;
                        } else {                            // course_name 已经存在
                            validateSettings.error.run(option, option.prompts.error.beUsed.replace("{1}", option.value));
                            note_name_state = false;
                        }
                    })
                }
                else {                                  // 相等则代表已经存在过,直接提示异常信息
                    validateSettings.error.run(option, option.prompts.error.beUsed.replace("{1}", option.value));
                    note_name_state = false;
                }
            }
            else {      // 代表已经验证通过,执行 succeed 的回调函数
                validateSettings.succeed.run(option);
            }
        }
    },
    note_key_words:function (option) {
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
        $("#note_name").ilearningValidate(validatePrompt.note_name, validateFunction.note_name, true);
        $("#note_key_words").ilearningValidate(validatePrompt.note_key_words, validateFunction.note_key_words, true);
        return validateFunction.FORM_submit(["#note_name","#note_key_words"]);
    }
});

$(function () {
    //默认下 course_name 框获得焦点
    setTimeout(function() {
        $("#note_name").get(0).focus();
    }, 0);

    // course_name 验证
    $("#note_name").ilearningValidate(validatePrompt.note_name, validateFunction.note_name);
    // course_type 验证
    $("#note_key_words").ilearningValidate(validatePrompt.note_key_words, validateFunction.note_key_words);

    //表单提交验证和服务器请求
    $("#registsubmit").click(function() {
        if(CKEDITOR.instances.editor.getData()==""){
            alert("内容不能为空！");
            return false;
        }

        var flag = validateFunction.FORM_validate();
        if (flag) {
            $(this).attr({"disabled":"disabled"}).attr({"value":"提交中,请稍等"});
            $.ajax({
                type: "POST",
                url: "/note/edit",
                data: {
                    "note_id":$("input[name='note_id']").val(),
                    "note_name":$("input[name='note_name']").val(),
                    "note_key_words":$("textarea[name='note_key_words']").val(),
                    "note_content":CKEDITOR.instances.editor.getData()
                },
                success: function(data) {
                    if (data.flag == true) {
                        alert("保存成功!");
                        sessionStorage.setItem("openMenuRefer","note");
                        window.location = "/note/list?filter_type=personal";
                    }else{
                        $("#note_key_words").removeClass().addClass("error").html(data.msg);
                    }
                }
            });
        }
    });
});

// 判断是否是表单回显
function isFormReload() {
    var note_id = $("input[name='note_id']").val();
    if(note_id != "" && note_id != undefined && note_id != null){
        return true;
    }
    return false;
}
