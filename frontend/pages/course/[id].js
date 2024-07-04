import Navbar from "../../app/components/navBar";
import CourseDetail from '../../app/components/CourseDetail';

export default function CoursePage() {
    return(
        <div className="relative">
            <Navbar/>
            <div className="mt-32">
                <CourseDetail/>
            </div>
        </div>
    );
}
