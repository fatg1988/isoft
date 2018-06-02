$(function () {
    // 以对象方式作为配置参数
    var ckeditor = CKEDITOR.replace('editor', {
        uiColor: '#ffccdd',
        height: '925',
        removePlugins :'elementspath,resize', // 移除编辑器底部状态栏显示的元素路径和调整编辑器大小的按钮
    });


    reloadNoteById(ckeditor);
});

// 表单回显
function reloadNoteById(ckeditor) {
    var note_id = $("input[name='note_id']").val();
    if(note_id != "" && note_id != undefined && note_id != null){
        $.ajax({
            url:"/note/queryNoteById",
            type:"post",
            data:{"note_id":note_id},
            success:function (data) {
                var note = data.note;
                $("input[name='note_name']").val(note.note_name);
                $("textarea[name='note_key_words']").text(note.note_key_words);
                ckeditor.setData(note.note_content);
            }
        });
    }
}