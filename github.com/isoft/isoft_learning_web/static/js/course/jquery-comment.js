function expand(currentNode) {
    if($(".comment_form").is(":hidden")){
        $(".comment_form").show();    //如果元素为隐藏,则将它显现
        $(currentNode).html("隐藏");
    }else{
        $(".comment_form").hide();     //如果元素为显现,则将其隐藏
        $(currentNode).html("展开");
    }
}



function showSubReply(currentNode,parent_id) {
    // 获取 topic_id 和 topic_type
    var topic_id = $(document).data("topic_id");
    var topic_type = $(document).data("topic_type");

    $.ajax({
        url:"/comment/topicReply/filter",
        method:"post",
        data:{"parent_id":parent_id,"topic_id":topic_id,"topic_type":topic_type},
        success:function (result) {
            if(result.status == "SUCCESS" && result.topic_replys.length > 0){
                if($("#sub_topic_reply_" + parent_id).is(":hidden")){
                    $("#sub_topic_reply_" + parent_id).show();    //如果元素为隐藏,则将它显现
                    $(currentNode).html("收起回复");

                    // 替换默认分隔符
                    $.views.settings.delimiters("[[", "]]");

                    $("#sub_topic_reply_" + parent_id).html("");

                    //获取模板
                    var jsRenderTpl = $.templates('#topic_reply_template');
                    //模板与数据结合
                    var finalTpl = jsRenderTpl({"topic_replys":result.topic_replys});
                    $("#sub_topic_reply_" + parent_id).html(finalTpl);


                }else{
                    $("#sub_topic_reply_" + parent_id).hide();     //如果元素为显现,则将其隐藏
                    $(currentNode).html("查看所有回复(" + result.topic_replys.length + ")");
                }
            }
        }
    });
}

function addSubReply(currentNode, refer_user_name, parent_id) {
    // 移动到评论锚点
    window.location.href="#reply_anchor";
    $(".comment_form").show();

    // 设置父评论 id
    $(document).data("_parent_id", parent_id);
    // 设置被评论人员
    $(document).data("_refer_user_name", refer_user_name);

    $("input[name='_refer_user_name_label']").val("回复@" + refer_user_name);
}

(function ($) {
    $.fn.commentComponent = function(options) {
        var defaults = {};
        options = $.extend(defaults, options);
        var methods = {
            reloadComment:function (parent_id, topic_id, topic_type) {
                $.ajax({
                    url:"/comment/topicReply/filter",
                    method:"post",
                    data:{"parent_id":parent_id,"topic_id":topic_id,"topic_type":topic_type},
                    success:function (result) {
                        if(result.status == "SUCCESS"){
                            // 替换默认分隔符
                            $.views.settings.delimiters("[[", "]]");

                            $("#topic_reply_area").html("");

                            //获取模板
                            var jsRenderTpl = $.templates('#topic_reply_template');
                            //模板与数据结合
                            var finalTpl = jsRenderTpl({"topic_replys":result.topic_replys});
                            $('#topic_reply_area').html(finalTpl);
                        }
                    }
                });
            },
            init_submit:function () {
                $("#submit_comment").click(function () {
                    // 获取父评论 id
                    var parent_id = $(document).data("_parent_id");
                    // 获取被评论人员
                    var refer_user_name = $(document).data("refer_user_name");
                    if(isEmpty(refer_user_name)){
                        // 隐藏域里面没有则使用初始化的值
                        refer_user_name = options.refer_user_name;
                    }
                    // 获取评论内容
                    var reply_content = $("textarea[name='reply_content']").val();

                    $.ajax({
                        url:"/comment/topicReply/add",
                        method:"post",
                        data:{"parent_id":parent_id,"reply_content":reply_content,"topic_id":options.topic_id,
                            "topic_type":options.topic_type,"refer_user_name":refer_user_name},
                        success:function (data) {
                            if(data.status == "SUCCESS"){
                                alert("提交成功!");
                                // 表单提交默认显示根级评论
                                methods.reloadComment(0, options.topic_id, options.topic_type);

                                // 初始化相关参数
                                methods.init_references();

                            }else{
                                alert("提交失败!");
                            }
                        }
                    });
                });
            },
            init_references:function (options) {
                // 初始化参数
                $(document).data("_parent_id", 0);
                $(document).data("_refer_user_name", options.refer_user_name);
                $(document).data("topic_id", options.topic_id);
                $(document).data("topic_type", options.topic_type);
            },
            init:function (options) {
                methods.init_references(options);

                // 渲染 topicTheme 信息
                $.ajax({
                    url:"/comment/topicTheme/filter",
                    method:"post",
                    data:{"topic_id":options.topic_id,"topic_type":options.topic_type},
                    async:false,
                    success:function (data) {
                        if(data.status == "SUCCESS"){
                            new Vue({
                                // 修改 vue 默认分隔符,解决冲突问题
                                delimiters: ['[[', ']]'],
                                el: '#topic_theme',
                                data: {
                                    topic_theme: data.topic_theme
                                }
                            });
                        }
                    }
                });

                // 初始化评论提交
                methods.init_submit();

                // 默认加载 parent_id = 0 根节点评论
                methods.reloadComment(0, options.topic_id, options.topic_type);
            }
        }

        methods.init(options);
        return this;
    }
})(jQuery);

// 使用范例
// $(function () {
//     // 获取 topic_id 和 topic_type
//     var topic_id = $("input[name='topic_id'][type='hidden']").val();
//     var refer_user_name = $("input[name='refer_user_name'][type='hidden']").val();
//     var topic_type = $("input[name='topic_type'][type='hidden']").val();
//
//     $("#comment_component").commentComponent({
//         "topic_id":topic_id,
//         "topic_type":topic_type,
//         "refer_user_name":refer_user_name,
//     });
// });

function isEmpty(obj){
    if(typeof obj == "undefined" || obj == null || obj == ""){
        return true;
    }else{
        return false;
    }
}