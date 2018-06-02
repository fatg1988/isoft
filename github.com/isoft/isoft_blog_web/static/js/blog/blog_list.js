$(function () {
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

    initPage();

    function initPage() {
        // 页面一初始化加载所有的类别信息
        loadAllBlog(1,10,true);
    }

    // 加载所有的类别信息
    function loadAllBlog(current_page,offset,renderPageTool) {
        // 加载内容
        $.ajax({
            url:"/blog/list",
            type:"post",
            data:{"current_page":current_page, "offset":offset, "personal":"personal"},
            success:function (obj) {
                var data = JSON.parse(obj);
                // 使用 $set 去修改这个 vueData 进行刷新页面
                blogVue.$set(blogVueData, 'blogs', data.blogs);
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

function delete_blog(blog_id) {
    if(window.confirm('确认要删除该博文吗?')){
        $.ajax({
            url:"/blog/delete",
            type:"post",
            data:{"blog_id":blog_id},
            success:function (data) {
                if(data.status == "SUCCESS"){
                    window.location.reload();
                }
            }
        });
    }
}

function edit_blog(blog_id) {
    window.location.href="/blog/edit?blog_id=" + blog_id;
}

function publish_blog(blog_id) {
    if(window.confirm('确认要发布该博文吗?发布后所有人可见!')){
        $.ajax({
            url:"/blog/publish",
            type:"post",
            data:{"blog_id":blog_id},
            success:function (data) {
                if(data.status == "SUCCESS"){
                    window.location.reload();
                }
            }
        });
    }
}
