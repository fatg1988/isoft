$(function () {
    // 定义一个全局的 vueData,初始数据为空
    var catalogVueData = {
        catalogs:[]
    };
    // 定义一个全局的 vue 实例,引用这个全局的 vueData
    var catalogVue = new Vue({
        // 修改 vue 默认分隔符,解决冲突问题
        delimiters: ['[[', ']]'],
        el: '#catalog_list',
        data: catalogVueData
    });

    initPage();

    function initPage() {
        // 页面一初始化加载所有的类别信息
        loadAllCatalog(1,10,true);
    }

    // 加载所有的类别信息
    function loadAllCatalog(current_page,offset,renderPageTool) {
        // 加载内容
        $.ajax({
            url:"/catalog/list",
            type:"post",
            data:{"current_page":current_page, "offset":offset, "personal":"personal"},
            success:function (obj) {
                var data = JSON.parse(obj);
                // 使用 $set 去修改这个 vueData 进行刷新页面
                catalogVue.$set(catalogVueData, 'catalogs', data.catalogs);
                if(renderPageTool == true){
                    // 渲染分页信息
                    $("#catalog_pageTool").Paging({pagesize: data.paginator.pagesize,count:data.paginator.totalcount,current:1,
                        callback:function(page,size,count){
                            loadAllCatalog(page, size, false);
                        }
                    });
                }
            }
        });
    }
});

function delete_catalog(catalog_id) {
    if(window.confirm('刪除博客分类会同时删除分类下的所有博客,确定要删除吗?')){
        $.ajax({
            url:"/catalog/delete",
            type:"post",
            data:{"catalog_id":catalog_id},
            success:function (data) {
                if(data.status == "SUCCESS"){
                    window.location.reload();
                }
            }
        });
    }
}

function edit_catalog(catalog_id) {
    window.location.href="/catalog/edit?catalog_id=" + catalog_id;
}