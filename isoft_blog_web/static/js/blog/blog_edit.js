$(function () {
    // 以对象方式作为配置参数
    var ckeditor = CKEDITOR.replace('editor', {
        uiColor: '#ffccdd',
        height: '925',
        removePlugins :'elementspath,resize', // 移除编辑器底部状态栏显示的元素路径和调整编辑器大小的按钮
    });
});

function edit_blog() {
    var blog_id = $("input[name='blog_id']").val();
    var blog_title = $("input[name='blog_title']").val();
    var key_words = $("input[name='key_words']").val();
    var catalog_id = $("select[name='catalog_id']").val();
    var blog_type = $("select[name='blog_type']").val();
    var blog_status = $("select[name='blog_status']").val();
    var content = CKEDITOR.instances.editor.getData();

    $.ajax({
        url:"/blog/edit",
        type:"post",
        data:function(){
            if(blog_id == "" || blog_id == null || blog_id == undefined){
                return {"blog_title":blog_title, "key_words":key_words, "catalog_id":catalog_id,
                    "blog_type":blog_type,"blog_status":blog_status,"content":content};
            }
            return {"blog_id":blog_id,"blog_title":blog_title, "key_words":key_words, "catalog_id":catalog_id,
                "blog_type":blog_type,"blog_status":blog_status,"content":content};
        }(),
        success:function (data) {
            if(data.status=="SUCCESS"){
                window.location.href="/blog/list";
            }else{
                alert("error!");
            }
        }
    });
}