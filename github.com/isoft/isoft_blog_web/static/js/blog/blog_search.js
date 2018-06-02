$(function () {
    initCatalogData();
    initBlogData();
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

    // 加载内容
    $.ajax({
        url:"/catalog/list",
        type:"post",
        data:{"current_page":1, "offset":20},
        success:function (obj) {
            var data = JSON.parse(obj);
            // 使用 $set 去修改这个 vueData 进行刷新页面
            catalogVue.$set(catalogVueData, 'catalogs', data.catalogs);
        }
    });
}

function initBlogData() {
    // 定义一个全局的 vueData,初始数据为空
    var blogVueData = {
        blogs:[]
    };
    // 定义一个全局的 vue 实例,引用这个全局的 vueData
    var blogVue = new Vue({
        // 修改 vue 默认分隔符,解决冲突问题
        delimiters: ['[[', ']]'],
        el: '#blog_list',
        data: blogVueData
    });

    // 页面一初始化加载第一页信息
    loadAllBlog(1,10,true);

    // 加载所有的类别信息
    function loadAllBlog(current_page,offset,renderPageTool) {
        var search_data={"current_page":current_page, "offset":offset};

        var blog_search = localStorage.getItem("blog_search");
        if(blog_search != "" && blog_search != null && blog_search != undefined){
            // 合并对象,扩展查询条件
            search_data = $.extend({}, search_data, JSON.parse(blog_search));
        }

        // 加载内容
        $.ajax({
            url:"/blog/list",
            type:"post",
            data:search_data,
            success:function (obj) {
                var data = JSON.parse(obj);
                // 使用 $set 去修改这个 vueData 进行刷新页面
                blogVue.$set(blogVueData, 'blogs', data.blogs);
                if(renderPageTool == true){
                    // 渲染分页信息
                    $("#blog_pageTool").Paging({pagesize: data.paginator.pagesize,count:data.paginator.totalcount,current:1,
                        callback:function(page,size,count){
                            loadAllCatalog(page, size, false);
                        }
                    });
                }
            }
        });
    }
}