function addCatalog() {
    var catalog_name = $("input[name='catalog_name']").val();
    var catalog_desc = $("input[name='catalog_desc']").val();
    $.ajax({
        url:"/catalog/add",
        method:"post",
        data:{"catalog_name":catalog_name,"catalog_desc":catalog_desc},
        success:function (data) {
            if(data.status=="SUCCESS"){
                window.location.href="/catalog/list"
            }else{
                alert(data.errorMsg);
            }
        }
    });
}

