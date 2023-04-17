#!/usr/bin/env node

const fs = require('fs');
const fastify = require('fastify')({
    logger: true
})

fastify.get('/v1/api/hello', async (request, reply) => {
    reply.code(200).send({ response: 'hello without api-key', success:true, data:'fastify node' })
})

fastify.get('/v1/api/healthy', async (request, reply) => {
    reply.code(200).send({ response: 'healthy', success:true, data:'fastify node' })
})

fastify.get('/v1/api/add/', {
    schema: {
        headers: {
            type: 'object',
            properties: {
                cookie: {
                    type: 'string',
                    pattern: '^token=public-key-\\d+$'
                }
            }
        },
        querystring: {
            type: 'object',
            required: ['x', 'y'],
            properties: {
                x: { type: 'integer' },
                y: { type: 'integer' }
            }
        },
        response: {
            200: {
                type: 'object',
                properties: {
                    success: { type: 'boolean' },
                    response: { type: 'string' },
                    data: { type: 'number' }
                }
            }
        }
    }
}, async (request, reply) => {
    const cookieHeader = request.headers['cookie'];
    if (!cookieHeader || cookieHeader !== 'token=public-key-123') {
        reply.status(400).send('malformed key fastify ');
        return;
    }

    const { x, y } = request.query;
    const result = x + y;
    reply.header('Content-Security-Policy',
        "default-src 'self'; font-src 'self'; img-src 'self'; script-src 'self'; style-src 'self'; frame-src 'self'")
    return { success: true, response: 'Success fastify', data: result };
});


fastify.get('/v1/api/div/', {
    schema: {
        querystring: {
            type: 'object',
            required: ['x', 'y'],
            properties: {
                x: { type: 'number' },
                y: { type: 'number' }
            }
        },
        response: {
            200: {
                type: 'object',
                properties: {
                    success: { type: 'boolean' },
                    response: { type: 'string' },
                    data: { type: 'number' }
                }
            }
        }
    }
}, async (request, reply) => {
    const { x, y } = request.query;
    if (y === 0) {
        reply.status(400).send({ success: false, response: 'Division by zero is not allowed', data: null });
        return;
    }
    const result = x / y;
    reply.header('Content-Security-Policy',
        "default-src 'self'; font-src 'self'; img-src 'self'; script-src 'self'; style-src 'self'; frame-src 'self'")
    return { success: true, response: 'Success fastify', data: result };
});

fastify.get('/v1/api/mul/', {
    schema: {
        querystring: {
            type: 'object',
            required: ['x', 'y'],
            properties: {
                x: { type: 'number' },
                y: { type: 'number' }
            }
        },
        response: {
            200: {
                type: 'object',
                properties: {
                    success: { type: 'boolean' },
                    response: { type: 'string' },
                    data: { type: 'number' }
                }
            },
            400: {
                type: 'object',
                properties: {
                    success: { type: 'boolean' },
                    response: { type: 'string' },
                    data: { type: 'null' }
                }
            }
        }
    }
}, async (request, reply) => {
    const { x, y } = request.query;
    const result = x * y;
    reply.header('Content-Security-Policy',
        "default-src 'self'; font-src 'self'; img-src 'self'; script-src 'self'; style-src 'self'; frame-src 'self'")
    reply.header('Content-Type', 'application/json')
    if (reply.statusCode >= 299) {
        return { success: false, response: 'False fastify', data: null };
    }
    return { success: true, response: 'Success fastify', data: result };
});

fastify.get('/v1/api/sub/', {
    schema: {
        querystring: {
            type: 'object',
            required: ['x', 'y'],
            properties: {
                x: { type: 'number' },
                y: { type: 'number' }
            }
        },
        response: {
            200: {
                type: 'object',
                properties: {
                    success: { type: 'boolean' },
                    response: { type: 'string' },
                    data: { type: 'number' }
                }
            }
        }
    }
}, async (request, reply) => {
    const { x, y } = request.query;
    const result = x - y;
    reply.header('Content-Security-Policy',
        "default-src 'self'; font-src 'self'; img-src 'self'; script-src 'self'; style-src 'self'; frame-src 'self'")
    return { success: true, response: 'Success fastify', data: result };
});


fastify.listen(process.env.PORT || 3001, '::', function (err, address) {
    if (err) {
        fastify.log.error(err)
        process.exit(1)
    }
    fastify.log.info('We are up')
});

console.log('Node.js web server at port 3001 is running..');
