import { NextRequest, NextResponse } from "next/server";
import { allBlogs } from "../handlers";

const request = async (request: NextRequest) => {
    const type = request.url.split('/').reverse()[0];
    const token = request.headers.get('Authorization') || '';
    if (type == 'allBlogs') {
        try {
            const {data, status, statusText} = await allBlogs(token);
            // console.log(data)
            if (status - 400 < 0) {
                return NextResponse.json({
                    blog: data,
                    status: status
                })
            } else {
                return NextResponse.json({
                    error: statusText,
                    status: status
                })
            }
        } catch (error) {
            console.log(error)
            return NextResponse.json({
                error: error,
                status: 404
            })
        }
    }
}

export {
    request as GET
}