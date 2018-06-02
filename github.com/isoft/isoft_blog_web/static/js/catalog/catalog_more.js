$(function () {
    initCatalogData();
});

function initCatalogData() {
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

    // 页面一初始化加载第一页信息
    loadAllCatalog(1,100,true);

    // 加载所有的类别信息
    function loadAllCatalog(current_page,offset,renderPageTool) {
        // 加载内容
        $.ajax({
            url:"/catalog/list",
            type:"post",
            data:{"current_page":current_page, "offset":offset},
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
}

