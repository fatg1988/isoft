$(function () {
    renderComment();
});

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