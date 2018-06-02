$(function () {
    renderStar();
    renderComment();

    $(".toggle_favorite").click(function () {
        var request_href = $(this).attr("request_href");
        var currentNode = this;
        $.ajax({
            url:request_href,
            method:"post",
            data:{},
            success:function () {
                var html = $(currentNode).html();
                if(html.indexOf("收藏") > 0){
                    if(html.indexOf("加入收藏") >= 0){
                        $(currentNode).html("取消收藏");
                    }else{
                        $(currentNode).html("加入收藏");
                    }
                }else{
                    if(html.indexOf("我要点赞") >= 0){
                        $(currentNode).html("取消赞");
                    }else{
                        $(currentNode).html("我要点赞");
                    }
                }
            }
        });
    });
});

// 渲染得分情况
function renderStar() {
    $(".star").raty({
        number: 5, // 多少个星星设置
        hints: ['冷门', '一般', '比较热门', '热门', '非常热门'],
        score: function(){      // 初始值设置
            return $(this).attr('score') / 2;
        },
        path: "/static/common/raty-2.8.0/lib/images",
        precision: true, //是否包含小数
        readOnly:true
    });
}

function renderComment() {
    // 获取 topic_id 和 topic_type
    var topic_id = $("input[name='topic_id'][type='hidden']").val();
    var refer_user_name = $("input[name='refer_user_name'][type='hidden']").val();
    var topic_type = $("input[name='topic_type'][type='hidden']").val();

    $("#comment_component").commentComponent({
        "topic_id":topic_id,
        "topic_type":topic_type,
        "refer_user_name":refer_user_name,
    });
};