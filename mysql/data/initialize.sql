CREATE TABLE image (
    /* id will be UUID V4 */
    id CHAR(36) NOT NULL,
    /* url field will be TEXT since the size of the url can be at least 2083 (IE) but we'll just use TEXT */
    url TEXT NOT NULL,
    /* og_name is the original file name we got from the user */ 
    og_name VARCHAR(255) NOT NULL,
    /* deleted states whether or not the image has been deleted */ 
    deleted BOOLEAN DEFAULT false,
    /* description is the text description of the image */ 
    description VARCHAR(1024) DEFAULT '',
    /* title is the display name the user wants to use for the image */ 
    title VARCHAR(255) DEFAULT '',
    PRIMARY KEY (id)
)