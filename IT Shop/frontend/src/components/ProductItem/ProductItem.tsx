import "./ProductItem.css"
import { Link } from "react-router-dom";
import { setSelectedIndex } from "../../data/selectedIndex";
import { useEffect, useState } from "react";
import { apiUrl, GetImageByProductID } from "../../services/http";
import { ImageInterface } from "../../Interfaces/IImage";

function ProductItem(props: { product: any; searchText: any; category: any}){

    const {product, searchText, category} = props;

    const [images, setImages] = useState<ImageInterface[]>([]);

    async function getImages(){
        let res = await GetImageByProductID(product.ID)
        if (res) {
            setImages(res);
        }
    }

    useEffect(()=> {
        getImages()
    }, [searchText, category])

    const imageUrl = images.length > 0 ? `${apiUrl}/${images[0].FilePath}` : '';

    const num = product.PricePerPiece.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 })
    
    return (
        <Link to="/Selected" onClick={()=> setSelectedIndex(product.ID-1)}> 
            <div className="container-product-item">
                <div className="img-box">
                    <img src={imageUrl} alt="" />
                </div>
                <div className="detail-box">
                    <div className="box1">
                        <div className="productname">
                            {product.ProductName}
                        </div>
                        <div className="brand">
                            brand: {product.Brand.Name}
                        </div>
                    </div>
                    <div className="price">
                        ฿{num}
                    </div>
                </div>
            </div>
        </Link>
    )
}

export default ProductItem;