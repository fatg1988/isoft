function edit_catalog() {
    var catalog_id = $("input[name='catalog_id']").val();
    var catalog_name = $("input[name='catalog_name']").val();
    var catalog_desc = $("input[name='catalog_desc']").val();
    $.ajax({
        url:"/catalog/edit",
        method:"post",
        data:function(){
            if(catalog_id == "" || catalog_id == null || catalog_id == undefined){
                return {"catalog_name":catalog_name,"catalog_desc":catalog_desc};
            }
            return {"catalog_id":catalog_id,"catalog_name":catalog_name,"catalog_desc":catalog_desc};
        }(),
        success:function (data) {
            if(data.status=="SUCCESS"){
                window.location.href="/catalog/list"
            }else{
                alert(data.errorMsg);
            }
        }
    });
}

